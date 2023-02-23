package yoredale

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞特尔SettleBarony struct {
	feud.BaseBarony
}

var BSettle塞特尔 feud.Barony = &塞特尔SettleBarony{}

func init() {
	f := BSettle塞特尔.(*塞特尔SettleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "settle",
		TitleName: "塞特尔",
		TitleCode: "b_settle",
	}
}
