package montbeliard

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莱尔LayreBarony struct {
	feud.BaseBarony
}

var BLayre莱尔 feud.Barony = &莱尔LayreBarony{}

func init() {
    f := BLayre莱尔.(*莱尔LayreBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "layre",
		TitleName: "莱尔",
		TitleCode: "b_layre",
	}
}
