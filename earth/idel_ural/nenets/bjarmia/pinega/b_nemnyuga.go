package pinega

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 涅姆纽加NemnyugaBarony struct {
	feud.BaseBarony
}

var BNemnyuga涅姆纽加 feud.Barony = &涅姆纽加NemnyugaBarony{}

func init() {
    f := BNemnyuga涅姆纽加.(*涅姆纽加NemnyugaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nemnyuga",
		TitleName: "涅姆纽加",
		TitleCode: "b_nemnyuga",
	}
}
