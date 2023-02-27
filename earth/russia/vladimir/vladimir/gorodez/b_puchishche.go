package gorodez

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普奇谢PuchishcheBarony struct {
	feud.BaseBarony
}

var BPuchishche普奇谢 feud.Barony = &普奇谢PuchishcheBarony{}

func init() {
    f := BPuchishche普奇谢.(*普奇谢PuchishcheBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "puchishche",
		TitleName: "普奇谢",
		TitleCode: "b_puchishche",
	}
}
