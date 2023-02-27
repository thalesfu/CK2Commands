package ephesos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊阿索斯IassosBarony struct {
	feud.BaseBarony
}

var BIassos伊阿索斯 feud.Barony = &伊阿索斯IassosBarony{}

func init() {
    f := BIassos伊阿索斯.(*伊阿索斯IassosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "iassos",
		TitleName: "伊阿索斯",
		TitleCode: "b_iassos",
	}
}
