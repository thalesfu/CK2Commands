package bern

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 下森UnterseenBarony struct {
	feud.BaseBarony
}

var BUnterseen下森 feud.Barony = &下森UnterseenBarony{}

func init() {
    f := BUnterseen下森.(*下森UnterseenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "unterseen",
		TitleName: "下森",
		TitleCode: "b_unterseen",
	}
}
