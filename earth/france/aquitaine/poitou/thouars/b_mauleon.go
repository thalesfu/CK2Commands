package thouars

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫莱翁MauleonBarony struct {
	feud.BaseBarony
}

var BMauleon莫莱翁 feud.Barony = &莫莱翁MauleonBarony{}

func init() {
    f := BMauleon莫莱翁.(*莫莱翁MauleonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mauleon",
		TitleName: "莫莱翁",
		TitleCode: "b_mauleon",
	}
}
