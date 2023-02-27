package nurnberg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 菲尔特FurthBarony struct {
	feud.BaseBarony
}

var BFurth菲尔特 feud.Barony = &菲尔特FurthBarony{}

func init() {
    f := BFurth菲尔特.(*菲尔特FurthBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "furth",
		TitleName: "菲尔特",
		TitleCode: "b_furth",
	}
}
