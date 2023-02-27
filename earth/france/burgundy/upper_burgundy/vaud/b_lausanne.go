package vaud

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 洛桑LausanneBarony struct {
	feud.BaseBarony
}

var BLausanne洛桑 feud.Barony = &洛桑LausanneBarony{}

func init() {
    f := BLausanne洛桑.(*洛桑LausanneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lausanne",
		TitleName: "洛桑",
		TitleCode: "b_lausanne",
	}
}
