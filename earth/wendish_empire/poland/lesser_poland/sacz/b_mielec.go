package sacz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅莱茨MielecBarony struct {
	feud.BaseBarony
}

var BMielec梅莱茨 feud.Barony = &梅莱茨MielecBarony{}

func init() {
    f := BMielec梅莱茨.(*梅莱茨MielecBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mielec",
		TitleName: "梅莱茨",
		TitleCode: "b_mielec",
	}
}
