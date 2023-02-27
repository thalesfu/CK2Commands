package kusinagara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波伐PavaBarony struct {
	feud.BaseBarony
}

var BPava波伐 feud.Barony = &波伐PavaBarony{}

func init() {
    f := BPava波伐.(*波伐PavaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pava",
		TitleName: "波伐",
		TitleCode: "b_pava",
	}
}
