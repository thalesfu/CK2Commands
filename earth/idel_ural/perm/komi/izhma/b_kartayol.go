package izhma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡尔塔约尔KartayolBarony struct {
	feud.BaseBarony
}

var BKartayol卡尔塔约尔 feud.Barony = &卡尔塔约尔KartayolBarony{}

func init() {
    f := BKartayol卡尔塔约尔.(*卡尔塔约尔KartayolBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kartayol",
		TitleName: "卡尔塔约尔",
		TitleCode: "b_kartayol",
	}
}
