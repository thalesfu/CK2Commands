package ulster

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡里克弗格斯CarrickfergusBarony struct {
	feud.BaseBarony
}

var BCarrickfergus卡里克弗格斯 feud.Barony = &卡里克弗格斯CarrickfergusBarony{}

func init() {
	f := BCarrickfergus卡里克弗格斯.(*卡里克弗格斯CarrickfergusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "carrickfergus",
		TitleName: "卡里克弗格斯",
		TitleCode: "b_carrickfergus",
	}
}
