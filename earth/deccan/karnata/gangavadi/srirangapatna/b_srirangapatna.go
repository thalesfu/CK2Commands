package srirangapatna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 室利浪伽钵多那SrirangapatnaBarony struct {
	feud.BaseBarony
}

var BSrirangapatna室利浪伽钵多那 feud.Barony = &室利浪伽钵多那SrirangapatnaBarony{}

func init() {
	f := BSrirangapatna室利浪伽钵多那.(*室利浪伽钵多那SrirangapatnaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "srirangapatna",
		TitleName: "室利浪伽钵多那",
		TitleCode: "b_srirangapatna",
	}
}
