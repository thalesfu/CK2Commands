package riga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 锡古尔达SiguldaBarony struct {
	feud.BaseBarony
}

var BSigulda锡古尔达 feud.Barony = &锡古尔达SiguldaBarony{}

func init() {
    f := BSigulda锡古尔达.(*锡古尔达SiguldaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sigulda",
		TitleName: "锡古尔达",
		TitleCode: "b_sigulda",
	}
}
