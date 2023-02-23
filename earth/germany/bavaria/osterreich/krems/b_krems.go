package krems

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克雷米斯KremsBarony struct {
	feud.BaseBarony
}

var BKrems克雷米斯 feud.Barony = &克雷米斯KremsBarony{}

func init() {
	f := BKrems克雷米斯.(*克雷米斯KremsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "krems",
		TitleName: "克雷米斯",
		TitleCode: "b_krems",
	}
}
