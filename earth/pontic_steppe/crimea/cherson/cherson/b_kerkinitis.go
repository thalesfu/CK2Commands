package cherson

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凯尔基尼提斯KerkinitisBarony struct {
	feud.BaseBarony
}

var BKerkinitis凯尔基尼提斯 feud.Barony = &凯尔基尼提斯KerkinitisBarony{}

func init() {
    f := BKerkinitis凯尔基尼提斯.(*凯尔基尼提斯KerkinitisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kerkinitis",
		TitleName: "凯尔基尼提斯",
		TitleCode: "b_kerkinitis",
	}
}
