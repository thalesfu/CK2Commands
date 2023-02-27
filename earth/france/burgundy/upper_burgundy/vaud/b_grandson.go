package vaud

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格朗松GrandsonBarony struct {
	feud.BaseBarony
}

var BGrandson格朗松 feud.Barony = &格朗松GrandsonBarony{}

func init() {
    f := BGrandson格朗松.(*格朗松GrandsonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "grandson",
		TitleName: "格朗松",
		TitleCode: "b_grandson",
	}
}
