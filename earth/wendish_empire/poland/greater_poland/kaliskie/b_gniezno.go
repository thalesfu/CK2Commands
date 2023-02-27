package kaliskie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格涅兹诺GnieznoBarony struct {
	feud.BaseBarony
}

var BGniezno格涅兹诺 feud.Barony = &格涅兹诺GnieznoBarony{}

func init() {
    f := BGniezno格涅兹诺.(*格涅兹诺GnieznoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gniezno",
		TitleName: "格涅兹诺",
		TitleCode: "b_gniezno",
	}
}
