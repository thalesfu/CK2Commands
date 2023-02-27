package gent

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 根特GentBarony struct {
	feud.BaseBarony
}

var BGent根特 feud.Barony = &根特GentBarony{}

func init() {
    f := BGent根特.(*根特GentBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gent",
		TitleName: "根特",
		TitleCode: "b_gent",
	}
}
