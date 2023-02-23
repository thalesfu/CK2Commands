package prusa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 提亚提拉ThyatiraBarony struct {
	feud.BaseBarony
}

var BThyatira提亚提拉 feud.Barony = &提亚提拉ThyatiraBarony{}

func init() {
	f := BThyatira提亚提拉.(*提亚提拉ThyatiraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "thyatira",
		TitleName: "提亚提拉",
		TitleCode: "b_thyatira",
	}
}
