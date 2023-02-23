package coimbra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩内拉PenelaBarony struct {
	feud.BaseBarony
}

var BPenela佩内拉 feud.Barony = &佩内拉PenelaBarony{}

func init() {
	f := BPenela佩内拉.(*佩内拉PenelaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "penela",
		TitleName: "佩内拉",
		TitleCode: "b_penela",
	}
}
