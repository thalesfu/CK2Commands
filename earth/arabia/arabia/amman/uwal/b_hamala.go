package uwal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈马拉HamalaBarony struct {
	feud.BaseBarony
}

var BHamala哈马拉 feud.Barony = &哈马拉HamalaBarony{}

func init() {
    f := BHamala哈马拉.(*哈马拉HamalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hamala",
		TitleName: "哈马拉",
		TitleCode: "b_hamala",
	}
}
