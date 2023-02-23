package torki

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏恰瓦SuceavaBarony struct {
	feud.BaseBarony
}

var BSuceava苏恰瓦 feud.Barony = &苏恰瓦SuceavaBarony{}

func init() {
	f := BSuceava苏恰瓦.(*苏恰瓦SuceavaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "suceava",
		TitleName: "苏恰瓦",
		TitleCode: "b_suceava",
	}
}
