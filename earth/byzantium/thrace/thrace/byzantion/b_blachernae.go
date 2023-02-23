package byzantion

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗拉赫尔奈BlachernaeBarony struct {
	feud.BaseBarony
}

var BBlachernae弗拉赫尔奈 feud.Barony = &弗拉赫尔奈BlachernaeBarony{}

func init() {
	f := BBlachernae弗拉赫尔奈.(*弗拉赫尔奈BlachernaeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "blachernae",
		TitleName: "弗拉赫尔奈",
		TitleCode: "b_blachernae",
	}
}
