package leix

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 敦纳梅斯DunamaseBarony struct {
	feud.BaseBarony
}

var BDunamase敦纳梅斯 feud.Barony = &敦纳梅斯DunamaseBarony{}

func init() {
    f := BDunamase敦纳梅斯.(*敦纳梅斯DunamaseBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dunamase",
		TitleName: "敦纳梅斯",
		TitleCode: "b_dunamase",
	}
}
