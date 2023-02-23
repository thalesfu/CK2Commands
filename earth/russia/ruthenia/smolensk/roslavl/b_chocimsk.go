package roslavl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍季姆斯克ChocimskBarony struct {
	feud.BaseBarony
}

var BChocimsk霍季姆斯克 feud.Barony = &霍季姆斯克ChocimskBarony{}

func init() {
	f := BChocimsk霍季姆斯克.(*霍季姆斯克ChocimskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chocimsk",
		TitleName: "霍季姆斯克",
		TitleCode: "b_chocimsk",
	}
}
