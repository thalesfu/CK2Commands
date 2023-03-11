package kara-kum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克孜勒阿尔瓦特Kyzyl_arvatBarony struct {
	feud.BaseBarony
}

var BKyzyl_arvat克孜勒阿尔瓦特 feud.Barony = &克孜勒阿尔瓦特Kyzyl_arvatBarony{}

func init() {
    f := BKyzyl_arvat克孜勒阿尔瓦特.(*克孜勒阿尔瓦特Kyzyl_arvatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kyzyl_arvat",
		TitleName: "克孜勒阿尔瓦特",
		TitleCode: "b_kyzyl_arvat",
	}
}
