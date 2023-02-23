package selenge

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博斯BotsiyBarony struct {
	feud.BaseBarony
}

var BBotsiy博斯 feud.Barony = &博斯BotsiyBarony{}

func init() {
	f := BBotsiy博斯.(*博斯BotsiyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "botsiy",
		TitleName: "博斯",
		TitleCode: "b_botsiy",
	}
}
