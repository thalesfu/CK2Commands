package zaysan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃尔昌卡VolchankaBarony struct {
	feud.BaseBarony
}

var BVolchanka沃尔昌卡 feud.Barony = &沃尔昌卡VolchankaBarony{}

func init() {
	f := BVolchanka沃尔昌卡.(*沃尔昌卡VolchankaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "volchanka",
		TitleName: "沃尔昌卡",
		TitleCode: "b_volchanka",
	}
}
