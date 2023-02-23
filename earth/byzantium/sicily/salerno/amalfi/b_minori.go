package amalfi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米诺丽MinoriBarony struct {
	feud.BaseBarony
}

var BMinori米诺丽 feud.Barony = &米诺丽MinoriBarony{}

func init() {
	f := BMinori米诺丽.(*米诺丽MinoriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "minori",
		TitleName: "米诺丽",
		TitleCode: "b_minori",
	}
}
