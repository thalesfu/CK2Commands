package mainling

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 登木DengmuBarony struct {
	feud.BaseBarony
}

var BDengmu登木 feud.Barony = &登木DengmuBarony{}

func init() {
    f := BDengmu登木.(*登木DengmuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dengmu",
		TitleName: "登木",
		TitleCode: "b_dengmu",
	}
}
