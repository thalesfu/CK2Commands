package korchev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 潘提卡派翁PanticapeaBarony struct {
	feud.BaseBarony
}

var BPanticapea潘提卡派翁 feud.Barony = &潘提卡派翁PanticapeaBarony{}

func init() {
    f := BPanticapea潘提卡派翁.(*潘提卡派翁PanticapeaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "panticapea",
		TitleName: "潘提卡派翁",
		TitleCode: "b_panticapea",
	}
}
