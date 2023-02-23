package forde

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比格斯塔BygstadBarony struct {
	feud.BaseBarony
}

var BBygstad比格斯塔 feud.Barony = &比格斯塔BygstadBarony{}

func init() {
	f := BBygstad比格斯塔.(*比格斯塔BygstadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bygstad",
		TitleName: "比格斯塔",
		TitleCode: "b_bygstad",
	}
}
