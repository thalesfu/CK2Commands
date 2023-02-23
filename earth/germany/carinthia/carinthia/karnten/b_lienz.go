package karnten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利恩茨LienzBarony struct {
	feud.BaseBarony
}

var BLienz利恩茨 feud.Barony = &利恩茨LienzBarony{}

func init() {
	f := BLienz利恩茨.(*利恩茨LienzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lienz",
		TitleName: "利恩茨",
		TitleCode: "b_lienz",
	}
}
