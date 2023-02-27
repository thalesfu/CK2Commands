package bacs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 宗博尔ZomborBarony struct {
	feud.BaseBarony
}

var BZombor宗博尔 feud.Barony = &宗博尔ZomborBarony{}

func init() {
    f := BZombor宗博尔.(*宗博尔ZomborBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zombor",
		TitleName: "宗博尔",
		TitleCode: "b_zombor",
	}
}
