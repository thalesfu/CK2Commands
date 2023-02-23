package metz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣于连SaintjulienBarony struct {
	feud.BaseBarony
}

var BSaintjulien圣于连 feud.Barony = &圣于连SaintjulienBarony{}

func init() {
	f := BSaintjulien圣于连.(*圣于连SaintjulienBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saintjulien",
		TitleName: "圣于连",
		TitleCode: "b_saintjulien",
	}
}
