package chauragarh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆罗曼BarmanBarony struct {
	feud.BaseBarony
}

var BBarman婆罗曼 feud.Barony = &婆罗曼BarmanBarony{}

func init() {
	f := BBarman婆罗曼.(*婆罗曼BarmanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "barman",
		TitleName: "婆罗曼",
		TitleCode: "b_barman",
	}
}
