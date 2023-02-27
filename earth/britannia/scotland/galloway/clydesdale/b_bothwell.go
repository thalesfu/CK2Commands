package clydesdale

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博斯韦尔BothwellBarony struct {
	feud.BaseBarony
}

var BBothwell博斯韦尔 feud.Barony = &博斯韦尔BothwellBarony{}

func init() {
    f := BBothwell博斯韦尔.(*博斯韦尔BothwellBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bothwell",
		TitleName: "博斯韦尔",
		TitleCode: "b_bothwell",
	}
}
