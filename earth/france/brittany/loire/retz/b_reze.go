package retz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 勒泽RezeBarony struct {
	feud.BaseBarony
}

var BReze勒泽 feud.Barony = &勒泽RezeBarony{}

func init() {
	f := BReze勒泽.(*勒泽RezeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "reze",
		TitleName: "勒泽",
		TitleCode: "b_reze",
	}
}
