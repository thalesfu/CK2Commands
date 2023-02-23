package mzab

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌尔格拉OuarglaBarony struct {
	feud.BaseBarony
}

var BOuargla乌尔格拉 feud.Barony = &乌尔格拉OuarglaBarony{}

func init() {
	f := BOuargla乌尔格拉.(*乌尔格拉OuarglaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ouargla",
		TitleName: "乌尔格拉",
		TitleCode: "b_ouargla",
	}
}
