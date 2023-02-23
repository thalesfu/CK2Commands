package lesbos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡洛尼KalloniBarony struct {
	feud.BaseBarony
}

var BKalloni卡洛尼 feud.Barony = &卡洛尼KalloniBarony{}

func init() {
	f := BKalloni卡洛尼.(*卡洛尼KalloniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kalloni",
		TitleName: "卡洛尼",
		TitleCode: "b_kalloni",
	}
}
