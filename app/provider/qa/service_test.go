package qa

import (
	"bbs/test"
	"fmt"
	"github.com/gohade/hade/framework/contract"
	"github.com/gohade/hade/framework/provider/cache"
	"github.com/gohade/hade/framework/provider/config"
	"github.com/gohade/hade/framework/provider/log"
	"github.com/gohade/hade/framework/provider/orm"
	"github.com/gohade/hade/framework/provider/redis"
	"testing"
)

func Test_QA(t *testing.T) {
	container := test.InitBaseContainer()
	container.Bind(&config.HadeConfigProvider{})
	container.Bind(&log.HadeLogServiceProvider{})
	container.Bind(&orm.GormProvider{})
	container.Bind(&redis.RedisProvider{})
	container.Bind(&cache.HadeCacheProvider{})

	ormService := container.MustMake(contract.ORMKey).(contract.ORMService)
	db, err := ormService.GetDB()
	if err != nil {
		t.Fatal(err)
	}
	if err := db.AutoMigrate(&Question{}, &Answer{}); err != nil {
		t.Fatal(err)
	}
	if err := db.Exec("truncate table questions").Error; err != nil {
		t.Fatal(err)
	}
	if err := db.Exec("truncate table answers").Error; err != nil {
		t.Fatal(err)
	}

	tmp, err := NewQaService(container)
	if err != nil {
		t.Fatal(err)
	}
	qaService := tmp.(Service)
	fmt.Println(qaService)
}
