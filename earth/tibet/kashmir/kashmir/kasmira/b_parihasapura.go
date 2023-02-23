package kasmira

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波利诃娑补罗ParihasapuraBarony struct {
	feud.BaseBarony
}

var BParihasapura波利诃娑补罗 feud.Barony = &波利诃娑补罗ParihasapuraBarony{}

func init() {
	f := BParihasapura波利诃娑补罗.(*波利诃娑补罗ParihasapuraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "parihasapura",
		TitleName: "波利诃娑补罗",
		TitleCode: "b_parihasapura",
	}
}
