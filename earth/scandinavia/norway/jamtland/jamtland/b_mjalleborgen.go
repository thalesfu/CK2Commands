package jamtland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 姆耶勒博尔根MjalleborgenBarony struct {
	feud.BaseBarony
}

var BMjalleborgen姆耶勒博尔根 feud.Barony = &姆耶勒博尔根MjalleborgenBarony{}

func init() {
	f := BMjalleborgen姆耶勒博尔根.(*姆耶勒博尔根MjalleborgenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mjalleborgen",
		TitleName: "姆耶勒博尔根",
		TitleCode: "b_mjalleborgen",
	}
}
