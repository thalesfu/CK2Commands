package livs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库尔拉玛KullamaaBarony struct {
	feud.BaseBarony
}

var BKullamaa库尔拉玛 feud.Barony = &库尔拉玛KullamaaBarony{}

func init() {
    f := BKullamaa库尔拉玛.(*库尔拉玛KullamaaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kullamaa",
		TitleName: "库尔拉玛",
		TitleCode: "b_kullamaa",
	}
}
