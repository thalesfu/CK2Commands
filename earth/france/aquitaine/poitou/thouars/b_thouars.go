package thouars

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图阿尔ThouarsBarony struct {
	feud.BaseBarony
}

var BThouars图阿尔 feud.Barony = &图阿尔ThouarsBarony{}

func init() {
    f := BThouars图阿尔.(*图阿尔ThouarsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "thouars",
		TitleName: "图阿尔",
		TitleCode: "b_thouars",
	}
}
