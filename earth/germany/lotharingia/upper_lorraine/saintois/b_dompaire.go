package saintois

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 栋派尔DompaireBarony struct {
	feud.BaseBarony
}

var BDompaire栋派尔 feud.Barony = &栋派尔DompaireBarony{}

func init() {
    f := BDompaire栋派尔.(*栋派尔DompaireBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dompaire",
		TitleName: "栋派尔",
		TitleCode: "b_dompaire",
	}
}
