package imeretia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 戈古提GegutiBarony struct {
	feud.BaseBarony
}

var BGeguti戈古提 feud.Barony = &戈古提GegutiBarony{}

func init() {
    f := BGeguti戈古提.(*戈古提GegutiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "geguti",
		TitleName: "戈古提",
		TitleCode: "b_geguti",
	}
}
