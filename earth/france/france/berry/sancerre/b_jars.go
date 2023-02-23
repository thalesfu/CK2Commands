package sancerre

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雅尔JarsBarony struct {
	feud.BaseBarony
}

var BJars雅尔 feud.Barony = &雅尔JarsBarony{}

func init() {
	f := BJars雅尔.(*雅尔JarsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jars",
		TitleName: "雅尔",
		TitleCode: "b_jars",
	}
}
