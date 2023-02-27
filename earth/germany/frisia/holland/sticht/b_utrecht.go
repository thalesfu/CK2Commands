package sticht

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌德勒支UtrechtBarony struct {
	feud.BaseBarony
}

var BUtrecht乌德勒支 feud.Barony = &乌德勒支UtrechtBarony{}

func init() {
    f := BUtrecht乌德勒支.(*乌德勒支UtrechtBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "utrecht",
		TitleName: "乌德勒支",
		TitleCode: "b_utrecht",
	}
}
