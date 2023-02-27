package bern

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比尔BielBarony struct {
	feud.BaseBarony
}

var BBiel比尔 feud.Barony = &比尔BielBarony{}

func init() {
    f := BBiel比尔.(*比尔BielBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "biel",
		TitleName: "比尔",
		TitleCode: "b_biel",
	}
}
