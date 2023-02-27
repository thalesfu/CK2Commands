package vexin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 昂德利AndelyBarony struct {
	feud.BaseBarony
}

var BAndely昂德利 feud.Barony = &昂德利AndelyBarony{}

func init() {
    f := BAndely昂德利.(*昂德利AndelyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "andely",
		TitleName: "昂德利",
		TitleCode: "b_andely",
	}
}
