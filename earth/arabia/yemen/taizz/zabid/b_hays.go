package zabid

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 海斯HaysBarony struct {
	feud.BaseBarony
}

var BHays海斯 feud.Barony = &海斯HaysBarony{}

func init() {
    f := BHays海斯.(*海斯HaysBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hays",
		TitleName: "海斯",
		TitleCode: "b_hays",
	}
}
