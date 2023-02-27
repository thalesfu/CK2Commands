package debul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 兰蒂希LandhiBarony struct {
	feud.BaseBarony
}

var BLandhi兰蒂希 feud.Barony = &兰蒂希LandhiBarony{}

func init() {
    f := BLandhi兰蒂希.(*兰蒂希LandhiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "landhi",
		TitleName: "兰蒂希",
		TitleCode: "b_landhi",
	}
}
