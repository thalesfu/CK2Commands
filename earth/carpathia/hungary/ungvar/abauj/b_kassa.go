package abauj

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 考绍KassaBarony struct {
	feud.BaseBarony
}

var BKassa考绍 feud.Barony = &考绍KassaBarony{}

func init() {
	f := BKassa考绍.(*考绍KassaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kassa",
		TitleName: "考绍",
		TitleCode: "b_kassa",
	}
}
