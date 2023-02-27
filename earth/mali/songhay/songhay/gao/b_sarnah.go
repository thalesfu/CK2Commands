package gao

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨尔纳SarnahBarony struct {
	feud.BaseBarony
}

var BSarnah萨尔纳 feud.Barony = &萨尔纳SarnahBarony{}

func init() {
    f := BSarnah萨尔纳.(*萨尔纳SarnahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sarnah",
		TitleName: "萨尔纳",
		TitleCode: "b_sarnah",
	}
}
