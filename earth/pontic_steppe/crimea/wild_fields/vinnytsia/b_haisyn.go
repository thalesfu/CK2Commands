package vinnytsia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 海辛HaisynBarony struct {
	feud.BaseBarony
}

var BHaisyn海辛 feud.Barony = &海辛HaisynBarony{}

func init() {
    f := BHaisyn海辛.(*海辛HaisynBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "haisyn",
		TitleName: "海辛",
		TitleCode: "b_haisyn",
	}
}
