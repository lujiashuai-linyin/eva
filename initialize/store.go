package initialize

import (
	"eva/biz/dal/store"
	"eva/biz/dal/store/private"
)

func InitStore() {
	store.CalendarStore = &private.CalendarStore{}
}
