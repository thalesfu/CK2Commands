package kunduz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 昆都士KunduzBarony struct {
	feud.BaseBarony
}

var BKunduz昆都士 feud.Barony = &昆都士KunduzBarony{}

func init() {
	f := BKunduz昆都士.(*昆都士KunduzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kunduz",
		TitleName: "昆都士",
		TitleCode: "b_kunduz",
	}
}
