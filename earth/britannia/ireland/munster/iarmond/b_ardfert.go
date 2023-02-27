package iarmond

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿德弗特ArdfertBarony struct {
	feud.BaseBarony
}

var BArdfert阿德弗特 feud.Barony = &阿德弗特ArdfertBarony{}

func init() {
    f := BArdfert阿德弗特.(*阿德弗特ArdfertBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ardfert",
		TitleName: "阿德弗特",
		TitleCode: "b_ardfert",
	}
}
