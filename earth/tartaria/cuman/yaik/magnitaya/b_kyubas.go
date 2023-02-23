package magnitaya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 丘巴斯KyubasBarony struct {
	feud.BaseBarony
}

var BKyubas丘巴斯 feud.Barony = &丘巴斯KyubasBarony{}

func init() {
	f := BKyubas丘巴斯.(*丘巴斯KyubasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kyubas",
		TitleName: "丘巴斯",
		TitleCode: "b_kyubas",
	}
}
