package kiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞尔内SerneBarony struct {
	feud.BaseBarony
}

var BSerne塞尔内 feud.Barony = &塞尔内SerneBarony{}

func init() {
	f := BSerne塞尔内.(*塞尔内SerneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "serne",
		TitleName: "塞尔内",
		TitleCode: "b_serne",
	}
}
