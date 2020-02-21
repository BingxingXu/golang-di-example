//go:generate goqueryset -in app.entity.go
package app

import "github.com/jinzhu/gorm"

// gen:qs
type App struct {
	gorm.Model
	Cover    string
	Download string
}

// 因为生成的代码缺少provider, 需要手动在这里加上
func ProviderAppRespsitory(db *gorm.DB) AppQuerySet{
	return AppQuerySet{db: db}
}
