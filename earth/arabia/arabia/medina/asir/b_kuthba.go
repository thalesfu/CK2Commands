package asir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 呼图白KuthbaBarony struct {
	feud.BaseBarony
}

var BKuthba呼图白 feud.Barony = &呼图白KuthbaBarony{}

func init() {
	f := BKuthba呼图白.(*呼图白KuthbaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kuthba",
		TitleName: "呼图白",
		TitleCode: "b_kuthba",
	}
}
