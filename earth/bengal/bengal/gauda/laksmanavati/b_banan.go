package laksmanavati

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 班安BananBarony struct {
	feud.BaseBarony
}

var BBanan班安 feud.Barony = &班安BananBarony{}

func init() {
    f := BBanan班安.(*班安BananBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "banan",
		TitleName: "班安",
		TitleCode: "b_banan",
	}
}
