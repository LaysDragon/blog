package main

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	ginzap "github.com/gin-contrib/zap"
	_ "github.com/lib/pq"
	prettyconsole "github.com/thessem/zap-prettyconsole"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"

	"github.com/LaysDragon/blog/apps/server/db/pgrepo"
	"github.com/LaysDragon/blog/apps/server/internal"
	"github.com/LaysDragon/blog/apps/server/usecase"
	"github.com/LaysDragon/blog/apps/server/web"
	"github.com/gin-gonic/gin"
)

func errorWrap[T any](val T, err error) func(string) (T, error) {
	return func(msg string) (T, error) {
		if err != nil {
			return val, fmt.Errorf(msg, err)
		}
		return val, err
	}
}

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
	router.Use(ginzap.Ginzap(log, time.RFC3339, true))
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

func main() {
	app := fx.New(
		fx.Provide(
			internal.LoadConfig,
			func() *zap.Logger {
				return prettyconsole.NewLogger(zap.DebugLevel)
			},
			func(log *zap.Logger) *zap.SugaredLogger {
				return log.Sugar()
			},
			func(config internal.Config) (boil.ContextExecutor, error) {
				return errorWrap(sql.Open(config.DBType, config.DataSourceName))("unable to connect to database, %w")
			},
			web.GetValidator,
			pgrepo.NewPost,
			pgrepo.NewAccount,
			usecase.NewPost,
			usecase.NewAccount,
			web.NewPostController,
			web.NewAccountController,
			NewServer,
		),
		//TODO: remove moduletrace and stacktrace field by custom encoder wrapper
		// https://stackoverflow.com/questions/73469128/hide-sensitive-fields-in-uber-zap-go
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			fxlogger := &fxevent.ZapLogger{Logger: log.Named("fx")}
			fxlogger.UseLogLevel(zap.InfoLevel)
			return fxlogger
		}),
		fx.Invoke(
			web.SetupValidation,
			web.SetupRouter,
			StartServer,
		),
	)

	app.Run()
}
