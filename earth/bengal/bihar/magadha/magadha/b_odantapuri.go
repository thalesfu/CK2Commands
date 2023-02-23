package magadha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 飞行寺OdantapuriBarony struct {
	feud.BaseBarony
}

var BOdantapuri飞行寺 feud.Barony = &飞行寺OdantapuriBarony{}

func init() {
	f := BOdantapuri飞行寺.(*飞行寺OdantapuriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "odantapuri",
		TitleName: "飞行寺",
		TitleCode: "b_odantapuri",
	}
}
