package casbin

import (
	"log"

	"github.com/casbin/casbin/v2"
	xormadapter "github.com/casbin/xorm-adapter/v3"
)

type CasbinConfig struct {
	Enforce *casbin.Enforcer
}

var Casbin *CasbinConfig

func Initialize(databaseUrl, modelPath string) *CasbinConfig {
	adapter, err := xormadapter.NewAdapter("postgres", databaseUrl, true)
	if err != nil {
		log.Fatalf("Failed to create enforcer: %v", err)
	}

	e, err := casbin.NewEnforcer(modelPath, adapter)
	if err != nil {
		log.Fatalf("Failed to create enforcer: %v", err)
	}

	err = e.LoadPolicy()
	if err != nil {
		log.Fatalf("Failed to load policies: %v", err)
	}

	Casbin = &CasbinConfig{Enforce: e}

	return Casbin
}
