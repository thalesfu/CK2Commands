package trondelag

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼达罗斯NidarosBarony struct {
	feud.BaseBarony
}

var BNidaros尼达罗斯 feud.Barony = &尼达罗斯NidarosBarony{}

func init() {
    f := BNidaros尼达罗斯.(*尼达罗斯NidarosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nidaros",
		TitleName: "尼达罗斯",
		TitleCode: "b_nidaros",
	}
}
