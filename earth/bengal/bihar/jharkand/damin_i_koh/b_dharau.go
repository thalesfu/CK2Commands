package damin_i_koh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 陀卢DharauBarony struct {
	feud.BaseBarony
}

var BDharau陀卢 feud.Barony = &陀卢DharauBarony{}

func init() {
    f := BDharau陀卢.(*陀卢DharauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dharau",
		TitleName: "陀卢",
		TitleCode: "b_dharau",
	}
}
