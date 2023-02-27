package uda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托罗姆ToromBarony struct {
	feud.BaseBarony
}

var BTorom托罗姆 feud.Barony = &托罗姆ToromBarony{}

func init() {
    f := BTorom托罗姆.(*托罗姆ToromBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "torom",
		TitleName: "托罗姆",
		TitleCode: "b_torom",
	}
}
