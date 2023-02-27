package vairata

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 都斯DhosiBarony struct {
	feud.BaseBarony
}

var BDhosi都斯 feud.Barony = &都斯DhosiBarony{}

func init() {
    f := BDhosi都斯.(*都斯DhosiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dhosi",
		TitleName: "都斯",
		TitleCode: "b_dhosi",
	}
}
