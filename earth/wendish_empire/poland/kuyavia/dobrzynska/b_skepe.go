package dobrzynska

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯肯佩SkepeBarony struct {
	feud.BaseBarony
}

var BSkepe斯肯佩 feud.Barony = &斯肯佩SkepeBarony{}

func init() {
    f := BSkepe斯肯佩.(*斯肯佩SkepeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "skepe",
		TitleName: "斯肯佩",
		TitleCode: "b_skepe",
	}
}
