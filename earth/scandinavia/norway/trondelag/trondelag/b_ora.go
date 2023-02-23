package trondelag

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 厄拉OraBarony struct {
	feud.BaseBarony
}

var BOra厄拉 feud.Barony = &厄拉OraBarony{}

func init() {
	f := BOra厄拉.(*厄拉OraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ora",
		TitleName: "厄拉",
		TitleCode: "b_ora",
	}
}
