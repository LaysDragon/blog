package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/LaysDragon/blog/apps/server/config"
	"github.com/LaysDragon/blog/apps/server/db"
	"github.com/LaysDragon/blog/apps/server/db/pgrepo"
	"github.com/LaysDragon/blog/apps/server/perm"
	"github.com/LaysDragon/blog/apps/server/usecase"
	"github.com/LaysDragon/blog/apps/server/web"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	prettyconsole "github.com/thessem/zap-prettyconsole"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func NewServer(log *zap.Logger) *gin.Engine {
	log = log.Named("gin")
	gin.DebugPrintFunc = func(format string, values ...interface{}) {
		msg := strings.TrimSpace(fmt.Sprintf(format, values...))
		if len(strings.Split(msg, "\n")) > 1 {
			log.Debug("", prettyconsole.FormattedString("msg", msg))
		} else {
			log.Debug(msg)
		}
	}
	router := gin.New()
	router.Use(ginzap.Ginzap(log, time.RFC3339, true), web.CtxLoggerMidleware(log))
	//TODO: use prettyconsole.FormattedString on panic recover stack trace
	router.Use(ginzap.RecoveryWithZap(log, true))
	return router
}

func StartServer(lc fx.Lifecycle, router *gin.Engine, log *zap.Logger) {
	log = log.Named("gin")

	httpSrv := &http.Server{Addr: ":8080", Handler: router}
	log.Info("Server setup complete")

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", httpSrv.Addr)
			if err != nil {
				return fmt.Errorf("failed to listen on %v: %w", httpSrv.Addr, err)
			}
			go httpSrv.Serve(ln)
			log.Sugar().Info("Server start at", httpSrv.Addr)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			err := httpSrv.Shutdown(ctx)
			log.Info("Server is stopped")
			return err

		},
	})
}

var LogModule = fx.Module("logger",
	fx.Provide(
		func() *zap.Logger {
			return prettyconsole.NewLogger(zap.DebugLevel)
		},
		func(log *zap.Logger) *zap.SugaredLogger {
			return log.Sugar()
		},
	),
	fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
		fxlogger := &fxevent.ZapLogger{Logger: log.Named("fx")}
		fxlogger.UseLogLevel(zap.InfoLevel)
		return fxlogger
	}),
)

func main() {
	app := fx.New(
		LogModule,
		fx.Provide(
			NewServer,
		),
		config.Module,
		db.Module,
		perm.Module,
		pgrepo.Module,
		usecase.Module,
		//TODO: remove moduletrace and stacktrace field by custom encoder wrapper
		// https://stackoverflow.com/questions/73469128/hide-sensitive-fields-in-uber-zap-go

		web.Module,
		fx.Invoke(
			db.InitDbData,
			StartServer,
		),
	)

	app.Run()
}
