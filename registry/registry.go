/*
registry.go
@package github.com/KensukeSakakibara/webamp4/registry
@author Kensuke Sakakibara
@since 2019.08.29
@copyright Copyright (c) 2019 Kensuke Sakakibara
@note コンストラクタ注入
*/
package registry

import (
	"github.com/KensukeSakakibara/webamp4/application/usecase"
	"github.com/KensukeSakakibara/webamp4/domain/model"
	"github.com/KensukeSakakibara/webamp4/domain/repository"
	"github.com/KensukeSakakibara/webamp4/infrastructure/config"
	"github.com/KensukeSakakibara/webamp4/infrastructure/persistence/session"
	"github.com/KensukeSakakibara/webamp4/infrastructure/persistence/table"
	"github.com/KensukeSakakibara/webamp4/presentation"
	"github.com/KensukeSakakibara/webamp4/presentation/router"
	"github.com/gin-gonic/contrib/sessions"
)

// Presentation layer
func DiMigration() presentation.MigrationInterface {
	return presentation.NewMigration(DiTable())
}

func DiInit() presentation.InitInterface {
	return presentation.NewInit(DiConfig(), DiUserModel())
}

func DiRouter() presentation.RouterInterface {
	return presentation.NewRouter(DiConfig(), DiRedisStore(), DiAppRouter(), DiApiRouter())
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

func DiAppRouter() router.AppRouterInterface {
	return router.NewAppRouter(DiCommonRouter(), DiUsecase(), DiAppIndexUsecase())
}

func DiApiRouter() router.ApiRouterInterface {
	return router.NewApiRouter(DiCommonRouter(), DiApiUsersUsecase())
}

// Application layer
func DiUsecase() usecase.UsecaseInterface {
	return usecase.NewUsecase(DiSessionRepository(), DiUserModel())
}

func DiAppIndexUsecase() usecase.AppIndexUsecaseInterface {
	return usecase.NewAppIndexUsecase(DiSessionRepository(), DiUserModel())
}

func DiApiUsersUsecase() usecase.ApiUsersUsecaseInterface {
	return usecase.NewApiUsersUsecase(DiSessionRepository(), DiUserModel())
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
