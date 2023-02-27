package carcassonne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉斯图尔LastoursBarony struct {
	feud.BaseBarony
}

var BLastours拉斯图尔 feud.Barony = &拉斯图尔LastoursBarony{}

func init() {
    f := BLastours拉斯图尔.(*拉斯图尔LastoursBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lastours",
		TitleName: "拉斯图尔",
		TitleCode: "b_lastours",
	}
}
