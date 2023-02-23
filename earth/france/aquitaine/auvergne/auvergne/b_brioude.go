package auvergne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布里乌德BrioudeBarony struct {
	feud.BaseBarony
}

var BBrioude布里乌德 feud.Barony = &布里乌德BrioudeBarony{}

func init() {
	f := BBrioude布里乌德.(*布里乌德BrioudeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "brioude",
		TitleName: "布里乌德",
		TitleCode: "b_brioude",
	}
}
