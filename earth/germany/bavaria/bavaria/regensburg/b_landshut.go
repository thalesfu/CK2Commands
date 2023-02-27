package regensburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 兰茨胡特LandshutBarony struct {
	feud.BaseBarony
}

var BLandshut兰茨胡特 feud.Barony = &兰茨胡特LandshutBarony{}

func init() {
    f := BLandshut兰茨胡特.(*兰茨胡特LandshutBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "landshut",
		TitleName: "兰茨胡特",
		TitleCode: "b_landshut",
	}
}
