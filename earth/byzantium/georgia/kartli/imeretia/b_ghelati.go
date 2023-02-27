package imeretia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 盖拉蒂GhelatiBarony struct {
	feud.BaseBarony
}

var BGhelati盖拉蒂 feud.Barony = &盖拉蒂GhelatiBarony{}

func init() {
    f := BGhelati盖拉蒂.(*盖拉蒂GhelatiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ghelati",
		TitleName: "盖拉蒂",
		TitleCode: "b_ghelati",
	}
}
