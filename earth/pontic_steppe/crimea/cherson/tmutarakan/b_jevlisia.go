package tmutarakan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 耶夫利西亚JevlisiaBarony struct {
	feud.BaseBarony
}

var BJevlisia耶夫利西亚 feud.Barony = &耶夫利西亚JevlisiaBarony{}

func init() {
    f := BJevlisia耶夫利西亚.(*耶夫利西亚JevlisiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jevlisia",
		TitleName: "耶夫利西亚",
		TitleCode: "b_jevlisia",
	}
}
