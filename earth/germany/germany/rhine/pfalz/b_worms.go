package pfalz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃尔姆斯WormsBarony struct {
	feud.BaseBarony
}

var BWorms沃尔姆斯 feud.Barony = &沃尔姆斯WormsBarony{}

func init() {
    f := BWorms沃尔姆斯.(*沃尔姆斯WormsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "worms",
		TitleName: "沃尔姆斯",
		TitleCode: "b_worms",
	}
}
