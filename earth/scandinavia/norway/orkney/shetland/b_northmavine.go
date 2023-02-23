package shetland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诺赫马因NorthmavineBarony struct {
	feud.BaseBarony
}

var BNorthmavine诺赫马因 feud.Barony = &诺赫马因NorthmavineBarony{}

func init() {
	f := BNorthmavine诺赫马因.(*诺赫马因NorthmavineBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "northmavine",
		TitleName: "诺赫马因",
		TitleCode: "b_northmavine",
	}
}
