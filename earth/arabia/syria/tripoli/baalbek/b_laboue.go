package baalbek

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莱卜沃LaboueBarony struct {
	feud.BaseBarony
}

var BLaboue莱卜沃 feud.Barony = &莱卜沃LaboueBarony{}

func init() {
	f := BLaboue莱卜沃.(*莱卜沃LaboueBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "laboue",
		TitleName: "莱卜沃",
		TitleCode: "b_laboue",
	}
}
