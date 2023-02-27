package irtek

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 捷普洛耶TeployeBarony struct {
	feud.BaseBarony
}

var BTeploye捷普洛耶 feud.Barony = &捷普洛耶TeployeBarony{}

func init() {
    f := BTeploye捷普洛耶.(*捷普洛耶TeployeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "teploye",
		TitleName: "捷普洛耶",
		TitleCode: "b_teploye",
	}
}
