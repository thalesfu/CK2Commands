package sundgau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 坦恩ThannBarony struct {
	feud.BaseBarony
}

var BThann坦恩 feud.Barony = &坦恩ThannBarony{}

func init() {
	f := BThann坦恩.(*坦恩ThannBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "thann",
		TitleName: "坦恩",
		TitleCode: "b_thann",
	}
}
