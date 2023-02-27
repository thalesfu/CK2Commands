package frisia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈灵根HarlingenBarony struct {
	feud.BaseBarony
}

var BHarlingen哈灵根 feud.Barony = &哈灵根HarlingenBarony{}

func init() {
    f := BHarlingen哈灵根.(*哈灵根HarlingenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "harlingen",
		TitleName: "哈灵根",
		TitleCode: "b_harlingen",
	}
}
