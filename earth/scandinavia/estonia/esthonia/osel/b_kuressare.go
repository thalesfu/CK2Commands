package osel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库雷萨雷KuressareBarony struct {
	feud.BaseBarony
}

var BKuressare库雷萨雷 feud.Barony = &库雷萨雷KuressareBarony{}

func init() {
    f := BKuressare库雷萨雷.(*库雷萨雷KuressareBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kuressare",
		TitleName: "库雷萨雷",
		TitleCode: "b_kuressare",
	}
}
