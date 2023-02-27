package koshma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥乌连捷尔OurenterBarony struct {
	feud.BaseBarony
}

var BOurenter奥乌连捷尔 feud.Barony = &奥乌连捷尔OurenterBarony{}

func init() {
    f := BOurenter奥乌连捷尔.(*奥乌连捷尔OurenterBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ourenter",
		TitleName: "奥乌连捷尔",
		TitleCode: "b_ourenter",
	}
}
