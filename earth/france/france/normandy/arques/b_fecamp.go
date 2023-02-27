package arques

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 费康FecampBarony struct {
	feud.BaseBarony
}

var BFecamp费康 feud.Barony = &费康FecampBarony{}

func init() {
    f := BFecamp费康.(*费康FecampBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fecamp",
		TitleName: "费康",
		TitleCode: "b_fecamp",
	}
}
