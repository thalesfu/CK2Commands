package mithila

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 须乔那SugaunaBarony struct {
	feud.BaseBarony
}

var BSugauna须乔那 feud.Barony = &须乔那SugaunaBarony{}

func init() {
    f := BSugauna须乔那.(*须乔那SugaunaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sugauna",
		TitleName: "须乔那",
		TitleCode: "b_sugauna",
	}
}
