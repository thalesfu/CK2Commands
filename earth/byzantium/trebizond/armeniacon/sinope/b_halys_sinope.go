package sinope

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈吕斯Halys_sinopeBarony struct {
	feud.BaseBarony
}

var BHalys_sinope哈吕斯 feud.Barony = &哈吕斯Halys_sinopeBarony{}

func init() {
    f := BHalys_sinope哈吕斯.(*哈吕斯Halys_sinopeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "halys_sinope",
		TitleName: "哈吕斯",
		TitleCode: "b_halys_sinope",
	}
}
