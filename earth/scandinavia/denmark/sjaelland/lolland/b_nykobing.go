package lolland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼克宾NykobingBarony struct {
	feud.BaseBarony
}

var BNykobing尼克宾 feud.Barony = &尼克宾NykobingBarony{}

func init() {
	f := BNykobing尼克宾.(*尼克宾NykobingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nykobing",
		TitleName: "尼克宾",
		TitleCode: "b_nykobing",
	}
}
