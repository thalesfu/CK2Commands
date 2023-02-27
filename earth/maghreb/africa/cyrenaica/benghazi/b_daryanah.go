package benghazi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 代尔亚奈DaryanahBarony struct {
	feud.BaseBarony
}

var BDaryanah代尔亚奈 feud.Barony = &代尔亚奈DaryanahBarony{}

func init() {
    f := BDaryanah代尔亚奈.(*代尔亚奈DaryanahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "daryanah",
		TitleName: "代尔亚奈",
		TitleCode: "b_daryanah",
	}
}
