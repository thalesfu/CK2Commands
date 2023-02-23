package cordoba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔科莱阿AlcoleaBarony struct {
	feud.BaseBarony
}

var BAlcolea阿尔科莱阿 feud.Barony = &阿尔科莱阿AlcoleaBarony{}

func init() {
	f := BAlcolea阿尔科莱阿.(*阿尔科莱阿AlcoleaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alcolea",
		TitleName: "阿尔科莱阿",
		TitleCode: "b_alcolea",
	}
}
