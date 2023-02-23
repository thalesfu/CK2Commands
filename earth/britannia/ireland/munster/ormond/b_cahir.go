package ormond

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凯尔CahirBarony struct {
	feud.BaseBarony
}

var BCahir凯尔 feud.Barony = &凯尔CahirBarony{}

func init() {
	f := BCahir凯尔.(*凯尔CahirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cahir",
		TitleName: "凯尔",
		TitleCode: "b_cahir",
	}
}
