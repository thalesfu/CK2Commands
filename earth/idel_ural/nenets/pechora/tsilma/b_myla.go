package tsilma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅拉MylaBarony struct {
	feud.BaseBarony
}

var BMyla梅拉 feud.Barony = &梅拉MylaBarony{}

func init() {
    f := BMyla梅拉.(*梅拉MylaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "myla",
		TitleName: "梅拉",
		TitleCode: "b_myla",
	}
}
