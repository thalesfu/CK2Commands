package quena

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希卜兰杰ShiblanjahBarony struct {
	feud.BaseBarony
}

var BShiblanjah希卜兰杰 feud.Barony = &希卜兰杰ShiblanjahBarony{}

func init() {
    f := BShiblanjah希卜兰杰.(*希卜兰杰ShiblanjahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shiblanjah",
		TitleName: "希卜兰杰",
		TitleCode: "b_shiblanjah",
	}
}
