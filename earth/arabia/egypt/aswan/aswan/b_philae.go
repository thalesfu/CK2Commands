package aswan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 菲莱PhilaeBarony struct {
	feud.BaseBarony
}

var BPhilae菲莱 feud.Barony = &菲莱PhilaeBarony{}

func init() {
    f := BPhilae菲莱.(*菲莱PhilaeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "philae",
		TitleName: "菲莱",
		TitleCode: "b_philae",
	}
}
