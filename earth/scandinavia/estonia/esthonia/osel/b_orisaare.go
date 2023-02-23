package osel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥里萨雷OrisaareBarony struct {
	feud.BaseBarony
}

var BOrisaare奥里萨雷 feud.Barony = &奥里萨雷OrisaareBarony{}

func init() {
	f := BOrisaare奥里萨雷.(*奥里萨雷OrisaareBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "orisaare",
		TitleName: "奥里萨雷",
		TitleCode: "b_orisaare",
	}
}
