package stezycka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 登布林DeblinBarony struct {
	feud.BaseBarony
}

var BDeblin登布林 feud.Barony = &登布林DeblinBarony{}

func init() {
    f := BDeblin登布林.(*登布林DeblinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "deblin",
		TitleName: "登布林",
		TitleCode: "b_deblin",
	}
}
