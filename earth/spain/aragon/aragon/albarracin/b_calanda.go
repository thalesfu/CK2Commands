package albarracin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡兰达CalandaBarony struct {
	feud.BaseBarony
}

var BCalanda卡兰达 feud.Barony = &卡兰达CalandaBarony{}

func init() {
    f := BCalanda卡兰达.(*卡兰达CalandaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "calanda",
		TitleName: "卡兰达",
		TitleCode: "b_calanda",
	}
}
