package sripuri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴斯纳BasnaBarony struct {
	feud.BaseBarony
}

var BBasna巴斯纳 feud.Barony = &巴斯纳BasnaBarony{}

func init() {
    f := BBasna巴斯纳.(*巴斯纳BasnaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "basna",
		TitleName: "巴斯纳",
		TitleCode: "b_basna",
	}
}
