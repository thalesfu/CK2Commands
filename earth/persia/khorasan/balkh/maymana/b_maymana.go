package maymana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅马内MaymanaBarony struct {
	feud.BaseBarony
}

var BMaymana梅马内 feud.Barony = &梅马内MaymanaBarony{}

func init() {
	f := BMaymana梅马内.(*梅马内MaymanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maymana",
		TitleName: "梅马内",
		TitleCode: "b_maymana",
	}
}
