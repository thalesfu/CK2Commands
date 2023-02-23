package selenge

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 者台DzhidaBarony struct {
	feud.BaseBarony
}

var BDzhida者台 feud.Barony = &者台DzhidaBarony{}

func init() {
	f := BDzhida者台.(*者台DzhidaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dzhida",
		TitleName: "者台",
		TitleCode: "b_dzhida",
	}
}
