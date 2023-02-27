package akhmim

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊赫明AkhmimBarony struct {
	feud.BaseBarony
}

var BAkhmim伊赫明 feud.Barony = &伊赫明AkhmimBarony{}

func init() {
    f := BAkhmim伊赫明.(*伊赫明AkhmimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "akhmim",
		TitleName: "伊赫明",
		TitleCode: "b_akhmim",
	}
}
