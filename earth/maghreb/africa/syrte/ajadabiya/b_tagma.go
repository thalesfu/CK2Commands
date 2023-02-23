package ajadabiya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泰格梅TagmaBarony struct {
	feud.BaseBarony
}

var BTagma泰格梅 feud.Barony = &泰格梅TagmaBarony{}

func init() {
	f := BTagma泰格梅.(*泰格梅TagmaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tagma",
		TitleName: "泰格梅",
		TitleCode: "b_tagma",
	}
}
