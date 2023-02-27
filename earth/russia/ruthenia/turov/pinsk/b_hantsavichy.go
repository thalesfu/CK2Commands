package pinsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 甘采维奇HantsavichyBarony struct {
	feud.BaseBarony
}

var BHantsavichy甘采维奇 feud.Barony = &甘采维奇HantsavichyBarony{}

func init() {
    f := BHantsavichy甘采维奇.(*甘采维奇HantsavichyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hantsavichy",
		TitleName: "甘采维奇",
		TitleCode: "b_hantsavichy",
	}
}
