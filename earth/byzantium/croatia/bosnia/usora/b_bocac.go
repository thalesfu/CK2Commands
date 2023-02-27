package usora

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博查茨BocacBarony struct {
	feud.BaseBarony
}

var BBocac博查茨 feud.Barony = &博查茨BocacBarony{}

func init() {
    f := BBocac博查茨.(*博查茨BocacBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bocac",
		TitleName: "博查茨",
		TitleCode: "b_bocac",
	}
}
