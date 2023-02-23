package eichstadt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希尔施贝格HirschbergBarony struct {
	feud.BaseBarony
}

var BHirschberg希尔施贝格 feud.Barony = &希尔施贝格HirschbergBarony{}

func init() {
	f := BHirschberg希尔施贝格.(*希尔施贝格HirschbergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hirschberg",
		TitleName: "希尔施贝格",
		TitleCode: "b_hirschberg",
	}
}
