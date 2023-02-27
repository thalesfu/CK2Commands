package pettau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩陶PettauBarony struct {
	feud.BaseBarony
}

var BPettau佩陶 feud.Barony = &佩陶PettauBarony{}

func init() {
    f := BPettau佩陶.(*佩陶PettauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pettau",
		TitleName: "佩陶",
		TitleCode: "b_pettau",
	}
}
