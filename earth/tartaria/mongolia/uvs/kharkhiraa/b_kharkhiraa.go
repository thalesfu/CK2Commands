package kharkhiraa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈尔齐拉KharkhiraaBarony struct {
	feud.BaseBarony
}

var BKharkhiraa哈尔齐拉 feud.Barony = &哈尔齐拉KharkhiraaBarony{}

func init() {
	f := BKharkhiraa哈尔齐拉.(*哈尔齐拉KharkhiraaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kharkhiraa",
		TitleName: "哈尔齐拉",
		TitleCode: "b_kharkhiraa",
	}
}
