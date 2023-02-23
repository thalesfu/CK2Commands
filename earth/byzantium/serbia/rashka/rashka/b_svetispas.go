package rashka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯维帖斯帕斯SvetispasBarony struct {
	feud.BaseBarony
}

var BSvetispas斯维帖斯帕斯 feud.Barony = &斯维帖斯帕斯SvetispasBarony{}

func init() {
	f := BSvetispas斯维帖斯帕斯.(*斯维帖斯帕斯SvetispasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "svetispas",
		TitleName: "斯维帖斯帕斯",
		TitleCode: "b_svetispas",
	}
}
