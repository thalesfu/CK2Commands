package badghis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库哈斯姆KuhsimBarony struct {
	feud.BaseBarony
}

var BKuhsim库哈斯姆 feud.Barony = &库哈斯姆KuhsimBarony{}

func init() {
	f := BKuhsim库哈斯姆.(*库哈斯姆KuhsimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kuhsim",
		TitleName: "库哈斯姆",
		TitleCode: "b_kuhsim",
	}
}
