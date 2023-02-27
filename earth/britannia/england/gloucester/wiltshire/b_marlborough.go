package wiltshire

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫尔伯勒MarlboroughBarony struct {
	feud.BaseBarony
}

var BMarlborough莫尔伯勒 feud.Barony = &莫尔伯勒MarlboroughBarony{}

func init() {
    f := BMarlborough莫尔伯勒.(*莫尔伯勒MarlboroughBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "marlborough",
		TitleName: "莫尔伯勒",
		TitleCode: "b_marlborough",
	}
}
