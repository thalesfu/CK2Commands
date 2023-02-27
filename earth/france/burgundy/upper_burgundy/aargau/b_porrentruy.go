package aargau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波朗特吕PorrentruyBarony struct {
	feud.BaseBarony
}

var BPorrentruy波朗特吕 feud.Barony = &波朗特吕PorrentruyBarony{}

func init() {
    f := BPorrentruy波朗特吕.(*波朗特吕PorrentruyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "porrentruy",
		TitleName: "波朗特吕",
		TitleCode: "b_porrentruy",
	}
}
