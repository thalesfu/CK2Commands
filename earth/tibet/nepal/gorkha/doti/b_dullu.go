package doti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杜鲁DulluBarony struct {
	feud.BaseBarony
}

var BDullu杜鲁 feud.Barony = &杜鲁DulluBarony{}

func init() {
    f := BDullu杜鲁.(*杜鲁DulluBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dullu",
		TitleName: "杜鲁",
		TitleCode: "b_dullu",
	}
}
