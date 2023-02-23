package rhodos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊阿吕索斯IalysosBarony struct {
	feud.BaseBarony
}

var BIalysos伊阿吕索斯 feud.Barony = &伊阿吕索斯IalysosBarony{}

func init() {
	f := BIalysos伊阿吕索斯.(*伊阿吕索斯IalysosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ialysos",
		TitleName: "伊阿吕索斯",
		TitleCode: "b_ialysos",
	}
}
