package cagliari

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊格莱西亚斯IglesiasBarony struct {
	feud.BaseBarony
}

var BIglesias伊格莱西亚斯 feud.Barony = &伊格莱西亚斯IglesiasBarony{}

func init() {
	f := BIglesias伊格莱西亚斯.(*伊格莱西亚斯IglesiasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "iglesias",
		TitleName: "伊格莱西亚斯",
		TitleCode: "b_iglesias",
	}
}
