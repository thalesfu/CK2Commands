package kharibta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙特努夫ShatnufBarony struct {
	feud.BaseBarony
}

var BShatnuf沙特努夫 feud.Barony = &沙特努夫ShatnufBarony{}

func init() {
	f := BShatnuf沙特努夫.(*沙特努夫ShatnufBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shatnuf",
		TitleName: "沙特努夫",
		TitleCode: "b_shatnuf",
	}
}
