package abakan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科穆纳尔KommunarBarony struct {
	feud.BaseBarony
}

var BKommunar科穆纳尔 feud.Barony = &科穆纳尔KommunarBarony{}

func init() {
    f := BKommunar科穆纳尔.(*科穆纳尔KommunarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kommunar",
		TitleName: "科穆纳尔",
		TitleCode: "b_kommunar",
	}
}
