package hesse

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马堡MarburgBarony struct {
	feud.BaseBarony
}

var BMarburg马堡 feud.Barony = &马堡MarburgBarony{}

func init() {
	f := BMarburg马堡.(*马堡MarburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "marburg",
		TitleName: "马堡",
		TitleCode: "b_marburg",
	}
}
