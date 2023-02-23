package leiningen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 翁施泰因UngsteinBarony struct {
	feud.BaseBarony
}

var BUngstein翁施泰因 feud.Barony = &翁施泰因UngsteinBarony{}

func init() {
	f := BUngstein翁施泰因.(*翁施泰因UngsteinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ungstein",
		TitleName: "翁施泰因",
		TitleCode: "b_ungstein",
	}
}
