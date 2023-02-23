package sripatha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 费罗兹补FirozpurBarony struct {
	feud.BaseBarony
}

var BFirozpur费罗兹补 feud.Barony = &费罗兹补FirozpurBarony{}

func init() {
	f := BFirozpur费罗兹补.(*费罗兹补FirozpurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "firozpur",
		TitleName: "费罗兹补",
		TitleCode: "b_firozpur",
	}
}
