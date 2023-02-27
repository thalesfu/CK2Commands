package lhasa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 八廓BarkhorBarony struct {
	feud.BaseBarony
}

var BBarkhor八廓 feud.Barony = &八廓BarkhorBarony{}

func init() {
    f := BBarkhor八廓.(*八廓BarkhorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "barkhor",
		TitleName: "八廓",
		TitleCode: "b_barkhor",
	}
}
