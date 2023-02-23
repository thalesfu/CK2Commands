package savoie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尚贝里ChamberyBarony struct {
	feud.BaseBarony
}

var BChambery尚贝里 feud.Barony = &尚贝里ChamberyBarony{}

func init() {
	f := BChambery尚贝里.(*尚贝里ChamberyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chambery",
		TitleName: "尚贝里",
		TitleCode: "b_chambery",
	}
}
