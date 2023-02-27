package delta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索伊斯XoisBarony struct {
	feud.BaseBarony
}

var BXois索伊斯 feud.Barony = &索伊斯XoisBarony{}

func init() {
    f := BXois索伊斯.(*索伊斯XoisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "xois",
		TitleName: "索伊斯",
		TitleCode: "b_xois",
	}
}
