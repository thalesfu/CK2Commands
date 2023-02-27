package lettigalians

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库凯诺伊斯KokenoisBarony struct {
	feud.BaseBarony
}

var BKokenois库凯诺伊斯 feud.Barony = &库凯诺伊斯KokenoisBarony{}

func init() {
    f := BKokenois库凯诺伊斯.(*库凯诺伊斯KokenoisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kokenois",
		TitleName: "库凯诺伊斯",
		TitleCode: "b_kokenois",
	}
}
