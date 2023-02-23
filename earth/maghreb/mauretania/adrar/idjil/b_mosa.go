package idjil

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫萨MosaBarony struct {
	feud.BaseBarony
}

var BMosa莫萨 feud.Barony = &莫萨MosaBarony{}

func init() {
	f := BMosa莫萨.(*莫萨MosaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mosa",
		TitleName: "莫萨",
		TitleCode: "b_mosa",
	}
}
