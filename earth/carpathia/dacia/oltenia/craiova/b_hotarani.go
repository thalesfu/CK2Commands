package craiova

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍特拉尼HotaraniBarony struct {
	feud.BaseBarony
}

var BHotarani霍特拉尼 feud.Barony = &霍特拉尼HotaraniBarony{}

func init() {
    f := BHotarani霍特拉尼.(*霍特拉尼HotaraniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hotarani",
		TitleName: "霍特拉尼",
		TitleCode: "b_hotarani",
	}
}
