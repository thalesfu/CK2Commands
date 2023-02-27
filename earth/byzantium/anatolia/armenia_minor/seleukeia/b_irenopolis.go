package seleukeia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊雷诺波利斯IrenopolisBarony struct {
	feud.BaseBarony
}

var BIrenopolis伊雷诺波利斯 feud.Barony = &伊雷诺波利斯IrenopolisBarony{}

func init() {
    f := BIrenopolis伊雷诺波利斯.(*伊雷诺波利斯IrenopolisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "irenopolis",
		TitleName: "伊雷诺波利斯",
		TitleCode: "b_irenopolis",
	}
}
