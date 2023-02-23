package jaunpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迦提城GadhipuriBarony struct {
	feud.BaseBarony
}

var BGadhipuri迦提城 feud.Barony = &迦提城GadhipuriBarony{}

func init() {
	f := BGadhipuri迦提城.(*迦提城GadhipuriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gadhipuri",
		TitleName: "迦提城",
		TitleCode: "b_gadhipuri",
	}
}
