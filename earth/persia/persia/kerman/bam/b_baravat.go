package bam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴拉瓦特BaravatBarony struct {
	feud.BaseBarony
}

var BBaravat巴拉瓦特 feud.Barony = &巴拉瓦特BaravatBarony{}

func init() {
	f := BBaravat巴拉瓦特.(*巴拉瓦特BaravatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baravat",
		TitleName: "巴拉瓦特",
		TitleCode: "b_baravat",
	}
}
