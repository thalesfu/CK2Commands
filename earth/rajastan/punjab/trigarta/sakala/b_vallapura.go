package sakala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 筏罗补罗VallapuraBarony struct {
	feud.BaseBarony
}

var BVallapura筏罗补罗 feud.Barony = &筏罗补罗VallapuraBarony{}

func init() {
	f := BVallapura筏罗补罗.(*筏罗补罗VallapuraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vallapura",
		TitleName: "筏罗补罗",
		TitleCode: "b_vallapura",
	}
}
