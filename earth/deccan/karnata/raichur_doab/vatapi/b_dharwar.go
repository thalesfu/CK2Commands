package vatapi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔尔瓦尔DharwarBarony struct {
	feud.BaseBarony
}

var BDharwar塔尔瓦尔 feud.Barony = &塔尔瓦尔DharwarBarony{}

func init() {
    f := BDharwar塔尔瓦尔.(*塔尔瓦尔DharwarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dharwar",
		TitleName: "塔尔瓦尔",
		TitleCode: "b_dharwar",
	}
}
