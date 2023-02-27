package konjikala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克孜勒阿尔瓦特KyzylarvatBarony struct {
	feud.BaseBarony
}

var BKyzylarvat克孜勒阿尔瓦特 feud.Barony = &克孜勒阿尔瓦特KyzylarvatBarony{}

func init() {
    f := BKyzylarvat克孜勒阿尔瓦特.(*克孜勒阿尔瓦特KyzylarvatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kyzylarvat",
		TitleName: "克孜勒阿尔瓦特",
		TitleCode: "b_kyzylarvat",
	}
}
