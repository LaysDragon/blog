package main

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/LaysDragon/blog/apps/server/db"
	"github.com/LaysDragon/blog/apps/server/db/pgrepo"
	"github.com/LaysDragon/blog/apps/server/internal"
	"github.com/LaysDragon/blog/apps/server/perm"
	"github.com/LaysDragon/blog/apps/server/usecase"
	"github.com/LaysDragon/blog/apps/server/web"
	"github.com/Thiht/transactor"
	stdlibTransactor "github.com/Thiht/transactor/stdlib"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	prettyconsole "github.com/thessem/zap-prettyconsole"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
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

func main() {
	//TODO: refactor with fx module
	app := fx.New(
		fx.Provide(
			internal.LoadConfig,
			func() *zap.Logger {
				return prettyconsole.NewLogger(zap.DebugLevel)
			},
			func(log *zap.Logger) *zap.SugaredLogger {
				return log.Sugar()
			},
			func(config internal.Config) (*sql.DB, error) {
				db, err := sql.Open(config.DBType, config.DataSourceName)
				if err != nil {
					return nil, fmt.Errorf("unable to connect to database, %w", err)
				}
				err = db.PingContext(context.Background())
				if err != nil {
					return nil, fmt.Errorf("unable to connect to database, %w", err)
				}
				return db, nil
			},

			func(db *sql.DB) boil.ContextExecutor {
				return db
			},
			func(db *sql.DB) (transactor.Transactor, stdlibTransactor.DBGetter) {
				return stdlibTransactor.NewTransactor(
					db,
					stdlibTransactor.NestedTransactionsSavepoints,
				)
			},

			func(config internal.Config, log *zap.Logger) *web.JwtHandler {
				return web.NewJwtHandler(config.JwtSecret, log)
			},
			func(config internal.Config, db *sql.DB, log *zap.Logger) (*perm.Perm, error) {
				return errorWrap(perm.New(db, config.DBType, log))("failed to init Permission service, %w")
			},
			web.GetValidator,
			pgrepo.NewPost,
			pgrepo.NewAccount,
			pgrepo.NewSite,
			pgrepo.NewSiteRole,
			usecase.NewPost,
			usecase.NewSite,
			usecase.NewAccount,
			web.NewPostController,
			web.NewAccountController,
			web.NewSiteController,
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
			db.InitDb,
			perm.InitPerm,
			db.InitDbData,
			web.SetupValidation,
			web.SetupRouter,
			StartServer,
		),
	)

	app.Run()
}
