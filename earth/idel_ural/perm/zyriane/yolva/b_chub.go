package yolva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 丘布ChubBarony struct {
	feud.BaseBarony
}

var BChub丘布 feud.Barony = &丘布ChubBarony{}

func init() {
    f := BChub丘布.(*丘布ChubBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chub",
		TitleName: "丘布",
		TitleCode: "b_chub",
	}
}
