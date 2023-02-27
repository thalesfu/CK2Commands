package marienburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃尔宾ElbingBarony struct {
	feud.BaseBarony
}

var BElbing埃尔宾 feud.Barony = &埃尔宾ElbingBarony{}

func init() {
    f := BElbing埃尔宾.(*埃尔宾ElbingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "elbing",
		TitleName: "埃尔宾",
		TitleCode: "b_elbing",
	}
}
