package osha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡尔加雷KargalyBarony struct {
	feud.BaseBarony
}

var BKargaly卡尔加雷 feud.Barony = &卡尔加雷KargalyBarony{}

func init() {
	f := BKargaly卡尔加雷.(*卡尔加雷KargalyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kargaly",
		TitleName: "卡尔加雷",
		TitleCode: "b_kargaly",
	}
}
