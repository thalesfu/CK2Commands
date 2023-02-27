package st_gallen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦杜兹VaduzBarony struct {
	feud.BaseBarony
}

var BVaduz瓦杜兹 feud.Barony = &瓦杜兹VaduzBarony{}

func init() {
    f := BVaduz瓦杜兹.(*瓦杜兹VaduzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vaduz",
		TitleName: "瓦杜兹",
		TitleCode: "b_vaduz",
	}
}
