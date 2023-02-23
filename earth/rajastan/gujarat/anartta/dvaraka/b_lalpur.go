package dvaraka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 那尔补罗LalpurBarony struct {
	feud.BaseBarony
}

var BLalpur那尔补罗 feud.Barony = &那尔补罗LalpurBarony{}

func init() {
	f := BLalpur那尔补罗.(*那尔补罗LalpurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lalpur",
		TitleName: "那尔补罗",
		TitleCode: "b_lalpur",
	}
}
