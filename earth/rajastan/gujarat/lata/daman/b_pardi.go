package daman

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波尔提PardiBarony struct {
	feud.BaseBarony
}

var BPardi波尔提 feud.Barony = &波尔提PardiBarony{}

func init() {
    f := BPardi波尔提.(*波尔提PardiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pardi",
		TitleName: "波尔提",
		TitleCode: "b_pardi",
	}
}
