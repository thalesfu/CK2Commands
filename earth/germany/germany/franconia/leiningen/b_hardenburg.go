package leiningen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈登堡HardenburgBarony struct {
	feud.BaseBarony
}

var BHardenburg哈登堡 feud.Barony = &哈登堡HardenburgBarony{}

func init() {
    f := BHardenburg哈登堡.(*哈登堡HardenburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hardenburg",
		TitleName: "哈登堡",
		TitleCode: "b_hardenburg",
	}
}
