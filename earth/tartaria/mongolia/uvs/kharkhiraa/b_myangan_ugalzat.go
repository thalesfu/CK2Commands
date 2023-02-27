package kharkhiraa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 明安乌嘎勒扎特Myangan_ugalzatBarony struct {
	feud.BaseBarony
}

var BMyangan_ugalzat明安乌嘎勒扎特 feud.Barony = &明安乌嘎勒扎特Myangan_ugalzatBarony{}

func init() {
    f := BMyangan_ugalzat明安乌嘎勒扎特.(*明安乌嘎勒扎特Myangan_ugalzatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "myangan_ugalzat",
		TitleName: "明安乌嘎勒扎特",
		TitleCode: "b_myangan_ugalzat",
	}
}
