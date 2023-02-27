package nagchu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达前DaqenBarony struct {
	feud.BaseBarony
}

var BDaqen达前 feud.Barony = &达前DaqenBarony{}

func init() {
    f := BDaqen达前.(*达前DaqenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "daqen",
		TitleName: "达前",
		TitleCode: "b_daqen",
	}
}
