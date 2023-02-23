package sinope

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 安德拉帕AndrapaBarony struct {
	feud.BaseBarony
}

var BAndrapa安德拉帕 feud.Barony = &安德拉帕AndrapaBarony{}

func init() {
	f := BAndrapa安德拉帕.(*安德拉帕AndrapaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "andrapa",
		TitleName: "安德拉帕",
		TitleCode: "b_andrapa",
	}
}
