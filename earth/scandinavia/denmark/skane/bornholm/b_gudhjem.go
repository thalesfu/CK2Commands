package bornholm

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古兹耶姆GudhjemBarony struct {
	feud.BaseBarony
}

var BGudhjem古兹耶姆 feud.Barony = &古兹耶姆GudhjemBarony{}

func init() {
	f := BGudhjem古兹耶姆.(*古兹耶姆GudhjemBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gudhjem",
		TitleName: "古兹耶姆",
		TitleCode: "b_gudhjem",
	}
}
