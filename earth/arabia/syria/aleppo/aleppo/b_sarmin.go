package aleppo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙敏SarminBarony struct {
	feud.BaseBarony
}

var BSarmin沙敏 feud.Barony = &沙敏SarminBarony{}

func init() {
    f := BSarmin沙敏.(*沙敏SarminBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sarmin",
		TitleName: "沙敏",
		TitleCode: "b_sarmin",
	}
}
