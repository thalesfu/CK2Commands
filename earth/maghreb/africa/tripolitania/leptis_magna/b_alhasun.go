package leptis_magna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾哈苏AlhasunBarony struct {
	feud.BaseBarony
}

var BAlhasun艾哈苏 feud.Barony = &艾哈苏AlhasunBarony{}

func init() {
    f := BAlhasun艾哈苏.(*艾哈苏AlhasunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alhasun",
		TitleName: "艾哈苏",
		TitleCode: "b_alhasun",
	}
}
