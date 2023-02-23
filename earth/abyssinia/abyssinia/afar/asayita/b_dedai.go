package asayita

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德代DedaiBarony struct {
	feud.BaseBarony
}

var BDedai德代 feud.Barony = &德代DedaiBarony{}

func init() {
	f := BDedai德代.(*德代DedaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dedai",
		TitleName: "德代",
		TitleCode: "b_dedai",
	}
}
