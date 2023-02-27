package vexin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 芒特MantesBarony struct {
	feud.BaseBarony
}

var BMantes芒特 feud.Barony = &芒特MantesBarony{}

func init() {
    f := BMantes芒特.(*芒特MantesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mantes",
		TitleName: "芒特",
		TitleCode: "b_mantes",
	}
}
