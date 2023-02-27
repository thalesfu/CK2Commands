package al_hasa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 穆拜赖兹AlmubarrazBarony struct {
	feud.BaseBarony
}

var BAlmubarraz穆拜赖兹 feud.Barony = &穆拜赖兹AlmubarrazBarony{}

func init() {
    f := BAlmubarraz穆拜赖兹.(*穆拜赖兹AlmubarrazBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "almubarraz",
		TitleName: "穆拜赖兹",
		TitleCode: "b_almubarraz",
	}
}
