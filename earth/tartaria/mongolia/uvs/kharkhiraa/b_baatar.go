package kharkhiraa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴特尔BaatarBarony struct {
	feud.BaseBarony
}

var BBaatar巴特尔 feud.Barony = &巴特尔BaatarBarony{}

func init() {
	f := BBaatar巴特尔.(*巴特尔BaatarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baatar",
		TitleName: "巴特尔",
		TitleCode: "b_baatar",
	}
}
