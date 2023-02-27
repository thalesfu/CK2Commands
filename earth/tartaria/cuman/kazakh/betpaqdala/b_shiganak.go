package betpaqdala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希加纳克ShiganakBarony struct {
	feud.BaseBarony
}

var BShiganak希加纳克 feud.Barony = &希加纳克ShiganakBarony{}

func init() {
    f := BShiganak希加纳克.(*希加纳克ShiganakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shiganak",
		TitleName: "希加纳克",
		TitleCode: "b_shiganak",
	}
}
