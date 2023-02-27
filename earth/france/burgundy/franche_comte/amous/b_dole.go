package amous

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多勒DoleBarony struct {
	feud.BaseBarony
}

var BDole多勒 feud.Barony = &多勒DoleBarony{}

func init() {
    f := BDole多勒.(*多勒DoleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dole",
		TitleName: "多勒",
		TitleCode: "b_dole",
	}
}
