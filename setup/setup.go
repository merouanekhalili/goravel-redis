package main

import (
	"os"

	"github.com/goravel/framework/packages"
	"github.com/goravel/framework/packages/match"
	"github.com/goravel/framework/packages/modify"
	"github.com/goravel/framework/support/path"
)

func main() {
	packages.Setup(os.Args).
		Install(
			modify.File(path.Config("app.go")).
				Find(match.Imports()).Modify(modify.AddImport(packages.GetModulePath())).
				Find(match.Providers()...).Modify(modify.AddProvider("&redis.ServiceProvider{}", "&cache.ServiceProvider{}")),
			modify.File(path.Config("cache.go")).
				Find(match.Imports()).Modify(modify.AddImport("github.com/goravel/framework/contracts/cache"), modify.AddImport("github.com/goravel/redis/facades", "redisfacades")).
				Find(match.Config("cache.stores")...).Modify(modify.AddConfig("redis", `map[string]any{ "driver": "custom", "connection": "default","via": func() (cache.Driver, error) {return redisfacades.Cache("redis")}}`)),
			modify.File(path.Config("queue.go")).
				Find(match.Imports()).Modify(modify.AddImport("github.com/goravel/framework/contracts/queue"), modify.AddImport("github.com/goravel/redis/facades", "redisfacades")).
				Find(match.Config("queue.connections")...).Modify(modify.AddConfig("redis", `map[string]any{ "driver": "custom", "connection": "default","via": func() (queue.Driver, error) {return redisfacades.Queue("redis")}}`)),
		).
		Uninstall(
			modify.File(path.Config("app.go")).
				Find(match.Imports()).Modify(modify.RemoveImport(packages.GetModulePath())).
				Find(match.Providers()...).Modify(modify.RemoveProvider("&redis.ServiceProvider{}")),
			modify.File(path.Config("cache.go")).
				Find(match.Imports()).Modify(modify.RemoveImport("github.com/goravel/framework/contracts/cache"), modify.RemoveImport("github.com/goravel/redis/facades", "redisfacades")).
				Find(match.Config("cache.stores")...).Modify(modify.RemoveConfig("redis")),
			modify.File(path.Config("queue.go")).
				Find(match.Imports()).Modify(modify.RemoveImport("github.com/goravel/framework/contracts/queue"), modify.RemoveImport("github.com/goravel/redis/facades", "redisfacades")).
				Find(match.Config("queue.connections")...).Modify(modify.RemoveConfig("redis")),
		).Execute()
}
