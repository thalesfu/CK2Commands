package amisos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 新凯撒利亚NeokaisareaBarony struct {
	feud.BaseBarony
}

var BNeokaisarea新凯撒利亚 feud.Barony = &新凯撒利亚NeokaisareaBarony{}

func init() {
	f := BNeokaisarea新凯撒利亚.(*新凯撒利亚NeokaisareaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "neokaisarea",
		TitleName: "新凯撒利亚",
		TitleCode: "b_neokaisarea",
	}
}
