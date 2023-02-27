package kaliskie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩兹德雷PyzdryBarony struct {
	feud.BaseBarony
}

var BPyzdry佩兹德雷 feud.Barony = &佩兹德雷PyzdryBarony{}

func init() {
    f := BPyzdry佩兹德雷.(*佩兹德雷PyzdryBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pyzdry",
		TitleName: "佩兹德雷",
		TitleCode: "b_pyzdry",
	}
}
