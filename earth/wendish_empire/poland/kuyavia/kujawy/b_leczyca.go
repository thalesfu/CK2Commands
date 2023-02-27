package kujawy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 文奇察LeczycaBarony struct {
	feud.BaseBarony
}

var BLeczyca文奇察 feud.Barony = &文奇察LeczycaBarony{}

func init() {
    f := BLeczyca文奇察.(*文奇察LeczycaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "leczyca",
		TitleName: "文奇察",
		TitleCode: "b_leczyca",
	}
}
