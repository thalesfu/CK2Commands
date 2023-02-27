package onega_peninsula

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 列特尼亚亚LetnyayaBarony struct {
	feud.BaseBarony
}

var BLetnyaya列特尼亚亚 feud.Barony = &列特尼亚亚LetnyayaBarony{}

func init() {
    f := BLetnyaya列特尼亚亚.(*列特尼亚亚LetnyayaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "letnyaya",
		TitleName: "列特尼亚亚",
		TitleCode: "b_letnyaya",
	}
}
