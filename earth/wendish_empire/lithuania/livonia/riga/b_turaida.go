package riga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图赖达TuraidaBarony struct {
	feud.BaseBarony
}

var BTuraida图赖达 feud.Barony = &图赖达TuraidaBarony{}

func init() {
    f := BTuraida图赖达.(*图赖达TuraidaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "turaida",
		TitleName: "图赖达",
		TitleCode: "b_turaida",
	}
}
