package dobrzynska

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利普诺LipnoBarony struct {
	feud.BaseBarony
}

var BLipno利普诺 feud.Barony = &利普诺LipnoBarony{}

func init() {
    f := BLipno利普诺.(*利普诺LipnoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lipno",
		TitleName: "利普诺",
		TitleCode: "b_lipno",
	}
}
