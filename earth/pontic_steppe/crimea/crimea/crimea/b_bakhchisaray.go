package crimea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴赫奇萨赖BakhchisarayBarony struct {
	feud.BaseBarony
}

var BBakhchisaray巴赫奇萨赖 feud.Barony = &巴赫奇萨赖BakhchisarayBarony{}

func init() {
    f := BBakhchisaray巴赫奇萨赖.(*巴赫奇萨赖BakhchisarayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bakhchisaray",
		TitleName: "巴赫奇萨赖",
		TitleCode: "b_bakhchisaray",
	}
}
