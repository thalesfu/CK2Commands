package seleukeia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科律科斯CorycusBarony struct {
	feud.BaseBarony
}

var BCorycus科律科斯 feud.Barony = &科律科斯CorycusBarony{}

func init() {
    f := BCorycus科律科斯.(*科律科斯CorycusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "corycus",
		TitleName: "科律科斯",
		TitleCode: "b_corycus",
	}
}
