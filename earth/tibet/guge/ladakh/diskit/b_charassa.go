package diskit

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 查拉萨CharassaBarony struct {
	feud.BaseBarony
}

var BCharassa查拉萨 feud.Barony = &查拉萨CharassaBarony{}

func init() {
	f := BCharassa查拉萨.(*查拉萨CharassaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "charassa",
		TitleName: "查拉萨",
		TitleCode: "b_charassa",
	}
}
