package vexin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉罗什_居永LarocheguyonBarony struct {
	feud.BaseBarony
}

var BLarocheguyon拉罗什_居永 feud.Barony = &拉罗什_居永LarocheguyonBarony{}

func init() {
	f := BLarocheguyon拉罗什_居永.(*拉罗什_居永LarocheguyonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "larocheguyon",
		TitleName: "拉罗什_居永",
		TitleCode: "b_larocheguyon",
	}
}
