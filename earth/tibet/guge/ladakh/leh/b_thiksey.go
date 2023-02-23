package leh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 提克西ThikseyBarony struct {
	feud.BaseBarony
}

var BThiksey提克西 feud.Barony = &提克西ThikseyBarony{}

func init() {
	f := BThiksey提克西.(*提克西ThikseyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "thiksey",
		TitleName: "提克西",
		TitleCode: "b_thiksey",
	}
}
