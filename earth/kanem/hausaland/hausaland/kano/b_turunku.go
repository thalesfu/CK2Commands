package kano

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图伦库TurunkuBarony struct {
	feud.BaseBarony
}

var BTurunku图伦库 feud.Barony = &图伦库TurunkuBarony{}

func init() {
	f := BTurunku图伦库.(*图伦库TurunkuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "turunku",
		TitleName: "图伦库",
		TitleCode: "b_turunku",
	}
}
