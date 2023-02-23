package dihistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克米尔KemirBarony struct {
	feud.BaseBarony
}

var BKemir克米尔 feud.Barony = &克米尔KemirBarony{}

func init() {
	f := BKemir克米尔.(*克米尔KemirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kemir",
		TitleName: "克米尔",
		TitleCode: "b_kemir",
	}
}
