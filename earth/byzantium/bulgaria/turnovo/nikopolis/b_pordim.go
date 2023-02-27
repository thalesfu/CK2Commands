package nikopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波尔迪姆PordimBarony struct {
	feud.BaseBarony
}

var BPordim波尔迪姆 feud.Barony = &波尔迪姆PordimBarony{}

func init() {
    f := BPordim波尔迪姆.(*波尔迪姆PordimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pordim",
		TitleName: "波尔迪姆",
		TitleCode: "b_pordim",
	}
}
