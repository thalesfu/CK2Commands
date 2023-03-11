package lutsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波恰耶夫PochaivBarony struct {
	feud.BaseBarony
}

var BPochaiv波恰耶夫 feud.Barony = &波恰耶夫PochaivBarony{}

func init() {
    f := BPochaiv波恰耶夫.(*波恰耶夫PochaivBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pochaiv",
		TitleName: "波恰耶夫",
		TitleCode: "b_pochaiv",
	}
}
