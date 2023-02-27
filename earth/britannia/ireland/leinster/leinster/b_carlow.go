package leinster

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡洛CarlowBarony struct {
	feud.BaseBarony
}

var BCarlow卡洛 feud.Barony = &卡洛CarlowBarony{}

func init() {
    f := BCarlow卡洛.(*卡洛CarlowBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "carlow",
		TitleName: "卡洛",
		TitleCode: "b_carlow",
	}
}
