package demetrias

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波多尼察BoudonitzaBarony struct {
	feud.BaseBarony
}

var BBoudonitza波多尼察 feud.Barony = &波多尼察BoudonitzaBarony{}

func init() {
	f := BBoudonitza波多尼察.(*波多尼察BoudonitzaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "boudonitza",
		TitleName: "波多尼察",
		TitleCode: "b_boudonitza",
	}
}
