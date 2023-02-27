package limbuwan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迈阿格朗德MyanglungBarony struct {
	feud.BaseBarony
}

var BMyanglung迈阿格朗德 feud.Barony = &迈阿格朗德MyanglungBarony{}

func init() {
    f := BMyanglung迈阿格朗德.(*迈阿格朗德MyanglungBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "myanglung",
		TitleName: "迈阿格朗德",
		TitleCode: "b_myanglung",
	}
}
