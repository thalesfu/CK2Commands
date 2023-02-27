package syrj

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩拉斯PyrasBarony struct {
	feud.BaseBarony
}

var BPyras佩拉斯 feud.Barony = &佩拉斯PyrasBarony{}

func init() {
    f := BPyras佩拉斯.(*佩拉斯PyrasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pyras",
		TitleName: "佩拉斯",
		TitleCode: "b_pyras",
	}
}
