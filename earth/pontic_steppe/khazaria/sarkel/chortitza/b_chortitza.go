package chortitza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍尔季察ChortitzaBarony struct {
	feud.BaseBarony
}

var BChortitza霍尔季察 feud.Barony = &霍尔季察ChortitzaBarony{}

func init() {
    f := BChortitza霍尔季察.(*霍尔季察ChortitzaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chortitza",
		TitleName: "霍尔季察",
		TitleCode: "b_chortitza",
	}
}
