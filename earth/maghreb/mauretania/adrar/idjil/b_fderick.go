package idjil

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗德里克FderickBarony struct {
	feud.BaseBarony
}

var BFderick弗德里克 feud.Barony = &弗德里克FderickBarony{}

func init() {
	f := BFderick弗德里克.(*弗德里克FderickBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fderick",
		TitleName: "弗德里克",
		TitleCode: "b_fderick",
	}
}
