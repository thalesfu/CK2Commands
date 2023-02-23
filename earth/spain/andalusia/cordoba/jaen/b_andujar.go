package jaen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 安杜哈尔AndujarBarony struct {
	feud.BaseBarony
}

var BAndujar安杜哈尔 feud.Barony = &安杜哈尔AndujarBarony{}

func init() {
	f := BAndujar安杜哈尔.(*安杜哈尔AndujarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "andujar",
		TitleName: "安杜哈尔",
		TitleCode: "b_andujar",
	}
}
