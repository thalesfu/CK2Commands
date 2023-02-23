package asir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泰巴莱TabalaBarony struct {
	feud.BaseBarony
}

var BTabala泰巴莱 feud.Barony = &泰巴莱TabalaBarony{}

func init() {
	f := BTabala泰巴莱.(*泰巴莱TabalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tabala",
		TitleName: "泰巴莱",
		TitleCode: "b_tabala",
	}
}
