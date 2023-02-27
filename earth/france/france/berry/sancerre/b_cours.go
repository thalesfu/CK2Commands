package sancerre

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库尔CoursBarony struct {
	feud.BaseBarony
}

var BCours库尔 feud.Barony = &库尔CoursBarony{}

func init() {
    f := BCours库尔.(*库尔CoursBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cours",
		TitleName: "库尔",
		TitleCode: "b_cours",
	}
}
