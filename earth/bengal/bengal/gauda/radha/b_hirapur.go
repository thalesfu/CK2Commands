package radha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 醯罗补罗HirapurBarony struct {
	feud.BaseBarony
}

var BHirapur醯罗补罗 feud.Barony = &醯罗补罗HirapurBarony{}

func init() {
	f := BHirapur醯罗补罗.(*醯罗补罗HirapurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hirapur",
		TitleName: "醯罗补罗",
		TitleCode: "b_hirapur",
	}
}
