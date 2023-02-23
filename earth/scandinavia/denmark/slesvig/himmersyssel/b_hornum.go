package himmersyssel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍努姆HornumBarony struct {
	feud.BaseBarony
}

var BHornum霍努姆 feud.Barony = &霍努姆HornumBarony{}

func init() {
	f := BHornum霍努姆.(*霍努姆HornumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hornum",
		TitleName: "霍努姆",
		TitleCode: "b_hornum",
	}
}
