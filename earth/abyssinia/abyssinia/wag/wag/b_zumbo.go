package wag

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 宗博ZumboBarony struct {
	feud.BaseBarony
}

var BZumbo宗博 feud.Barony = &宗博ZumboBarony{}

func init() {
    f := BZumbo宗博.(*宗博ZumboBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zumbo",
		TitleName: "宗博",
		TitleCode: "b_zumbo",
	}
}
