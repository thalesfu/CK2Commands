package vengipura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾鲁卢EluruBarony struct {
	feud.BaseBarony
}

var BEluru艾鲁卢 feud.Barony = &艾鲁卢EluruBarony{}

func init() {
	f := BEluru艾鲁卢.(*艾鲁卢EluruBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "eluru",
		TitleName: "艾鲁卢",
		TitleCode: "b_eluru",
	}
}
