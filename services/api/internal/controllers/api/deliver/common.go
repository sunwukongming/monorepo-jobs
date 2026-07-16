package deliver

import "app/models/bolejiang"

type DeliverPassage struct {
	Deliver bolejiang.Deliver `xorm:"extends"`
	Passage bolejiang.Passage `xorm:"extends"`
	Account bolejiang.Account `xorm:"extends"`
}
