package kamarupanagara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩尼矩吒ManikutaBarony struct {
	feud.BaseBarony
}

var BManikuta摩尼矩吒 feud.Barony = &摩尼矩吒ManikutaBarony{}

func init() {
	f := BManikuta摩尼矩吒.(*摩尼矩吒ManikutaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "manikuta",
		TitleName: "摩尼矩吒",
		TitleCode: "b_manikuta",
	}
}
