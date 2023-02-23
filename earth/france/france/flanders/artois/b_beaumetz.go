package artois

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博梅斯BeaumetzBarony struct {
	feud.BaseBarony
}

var BBeaumetz博梅斯 feud.Barony = &博梅斯BeaumetzBarony{}

func init() {
	f := BBeaumetz博梅斯.(*博梅斯BeaumetzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "beaumetz",
		TitleName: "博梅斯",
		TitleCode: "b_beaumetz",
	}
}
