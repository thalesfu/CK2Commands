package regensburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 因戈尔施塔特IngolstadtBarony struct {
	feud.BaseBarony
}

var BIngolstadt因戈尔施塔特 feud.Barony = &因戈尔施塔特IngolstadtBarony{}

func init() {
    f := BIngolstadt因戈尔施塔特.(*因戈尔施塔特IngolstadtBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ingolstadt",
		TitleName: "因戈尔施塔特",
		TitleCode: "b_ingolstadt",
	}
}
