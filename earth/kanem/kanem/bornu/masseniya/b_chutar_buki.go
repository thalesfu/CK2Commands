package masseniya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 楚塔布基Chutar_bukiBarony struct {
	feud.BaseBarony
}

var BChutar_buki楚塔布基 feud.Barony = &楚塔布基Chutar_bukiBarony{}

func init() {
    f := BChutar_buki楚塔布基.(*楚塔布基Chutar_bukiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chutar_buki",
		TitleName: "楚塔布基",
		TitleCode: "b_chutar_buki",
	}
}
