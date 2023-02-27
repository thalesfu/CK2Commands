package kanj_rustaq

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 费内斯克FeneskBarony struct {
	feud.BaseBarony
}

var BFenesk费内斯克 feud.Barony = &费内斯克FeneskBarony{}

func init() {
    f := BFenesk费内斯克.(*费内斯克FeneskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fenesk",
		TitleName: "费内斯克",
		TitleCode: "b_fenesk",
	}
}
