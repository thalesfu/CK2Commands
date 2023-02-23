package figuig

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德拉布利亚DrabliaBarony struct {
	feud.BaseBarony
}

var BDrablia德拉布利亚 feud.Barony = &德拉布利亚DrabliaBarony{}

func init() {
	f := BDrablia德拉布利亚.(*德拉布利亚DrabliaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "drablia",
		TitleName: "德拉布利亚",
		TitleCode: "b_drablia",
	}
}
