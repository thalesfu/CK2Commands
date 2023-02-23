package osel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莱西LeisiBarony struct {
	feud.BaseBarony
}

var BLeisi莱西 feud.Barony = &莱西LeisiBarony{}

func init() {
	f := BLeisi莱西.(*莱西LeisiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "leisi",
		TitleName: "莱西",
		TitleCode: "b_leisi",
	}
}
