package vodi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩赫基内林纳NoteborgBarony struct {
	feud.BaseBarony
}

var BNoteborg佩赫基内林纳 feud.Barony = &佩赫基内林纳NoteborgBarony{}

func init() {
    f := BNoteborg佩赫基内林纳.(*佩赫基内林纳NoteborgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "noteborg",
		TitleName: "佩赫基内林纳",
		TitleCode: "b_noteborg",
	}
}
