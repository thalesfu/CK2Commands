package gowrie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 福蒂维尔特ForteviotBarony struct {
	feud.BaseBarony
}

var BForteviot福蒂维尔特 feud.Barony = &福蒂维尔特ForteviotBarony{}

func init() {
    f := BForteviot福蒂维尔特.(*福蒂维尔特ForteviotBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "forteviot",
		TitleName: "福蒂维尔特",
		TitleCode: "b_forteviot",
	}
}
