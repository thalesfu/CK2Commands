package vladimir_volynsky

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫恰尔基MochalkyBarony struct {
	feud.BaseBarony
}

var BMochalky莫恰尔基 feud.Barony = &莫恰尔基MochalkyBarony{}

func init() {
    f := BMochalky莫恰尔基.(*莫恰尔基MochalkyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mochalky",
		TitleName: "莫恰尔基",
		TitleCode: "b_mochalky",
	}
}
