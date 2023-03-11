package kyzyl-su

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克纳尔KenarBarony struct {
	feud.BaseBarony
}

var BKenar克纳尔 feud.Barony = &克纳尔KenarBarony{}

func init() {
    f := BKenar克纳尔.(*克纳尔KenarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kenar",
		TitleName: "克纳尔",
		TitleCode: "b_kenar",
	}
}
