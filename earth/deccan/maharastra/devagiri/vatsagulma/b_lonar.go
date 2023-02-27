package vatsagulma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卢那罗LonarBarony struct {
	feud.BaseBarony
}

var BLonar卢那罗 feud.Barony = &卢那罗LonarBarony{}

func init() {
    f := BLonar卢那罗.(*卢那罗LonarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lonar",
		TitleName: "卢那罗",
		TitleCode: "b_lonar",
	}
}
