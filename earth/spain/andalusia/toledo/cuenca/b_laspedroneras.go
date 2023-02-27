package cuenca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉斯佩德罗涅拉斯LaspedronerasBarony struct {
	feud.BaseBarony
}

var BLaspedroneras拉斯佩德罗涅拉斯 feud.Barony = &拉斯佩德罗涅拉斯LaspedronerasBarony{}

func init() {
    f := BLaspedroneras拉斯佩德罗涅拉斯.(*拉斯佩德罗涅拉斯LaspedronerasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "laspedroneras",
		TitleName: "拉斯佩德罗涅拉斯",
		TitleCode: "b_laspedroneras",
	}
}
