/*
registry.go
@package github.com/KensukeSakakibara/gin_gorm_skeleton/registry
@author Kensuke Sakakibara
@since 2019.08.29
@copyright Copyright (c) 2019 Kensuke Sakakibara
@note コンストラクタ注入
*/
package registry

import (
	"github.com/KensukeSakakibara/gin_gorm_skeleton/application/usecase"
	"github.com/KensukeSakakibara/gin_gorm_skeleton/domain/model"
	"github.com/KensukeSakakibara/gin_gorm_skeleton/domain/repository"
	"github.com/KensukeSakakibara/gin_gorm_skeleton/infrastructure/config"
	"github.com/KensukeSakakibara/gin_gorm_skeleton/infrastructure/persistence/session"
	"github.com/KensukeSakakibara/gin_gorm_skeleton/infrastructure/persistence/table"
	"github.com/KensukeSakakibara/gin_gorm_skeleton/presentation/app"
	"github.com/KensukeSakakibara/gin_gorm_skeleton/presentation/app/router"
	"github.com/gin-gonic/contrib/sessions"
)

// Presentation layer
func DiMigration() app.MigrationInterface {
	return app.NewMigration(DiTable())
}

func DiInit() app.InitInterface {
	return app.NewInit(DiConfig(), DiUserModel())
}

func DiRouter() router.RouterInterface {
	return router.NewRouter(DiConfig(), DiRedisStore(), DiIndexRouter())
}

const REDIS_MAX_NUM = 10 // タイムアウトの秒数

func DiRedisStore() *sessions.RedisStore {
	configInstance := DiConfig()
	redisHost := configInstance.App.RedisHost
	redisPassword := configInstance.App.RedisPassword
	store, _ := sessions.NewRedisStore(REDIS_MAX_NUM, "tcp", redisHost, redisPassword, []byte("secret"))
	return &store
}

func DiCommonRouter() router.CommonRouterInterface {
	return router.NewCommonRouter(DiTableRepository())
}

func DiIndexRouter() router.IndexRouterInterface {
	return router.NewIndexRouter(DiCommonRouter(), DiUsecase(), DiIndexUsecase())
}

// Application layer
func DiUsecase() usecase.UsecaseInterface {
	return usecase.NewUsecase(DiSessionRepository(), DiUserModel())
}

func DiIndexUsecase() usecase.IndexUsecaseInterface {
	return usecase.NewIndexUsecase(DiSessionRepository(), DiUserModel())
}

// Domain layer
func DiUserModel() model.UserModelInterface {
	return model.NewUserModel(DiTableRepository(), DiUserRepository())
}

// Repository
func DiSessionRepository() repository.SessionRepositoryInterface {
	return repository.NewSessionRepository(DiSession())
}

func DiTableRepository() repository.TableRepositoryInterface {
	return repository.NewTableRepository(DiTable())
}

func DiUserRepository() repository.UserRepositoryInterface {
	return repository.NewUserRepository(DiTUsersTable())
}

// Infrastructure layer
func DiConfig() *config.Config {
	return config.NewConfig()
}

func DiSession() session.SessionInterface {
	return session.NewSession()
}

func DiTable() table.TableInterface {
	return table.NewTable(DiConfig())
}

func DiTUsersTable() table.TUsersTableInterface {
	return table.NewTUsersTable(DiTable())
}
