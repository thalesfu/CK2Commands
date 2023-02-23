package nandurbar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏坦普尔SultanpurBarony struct {
	feud.BaseBarony
}

var BSultanpur苏坦普尔 feud.Barony = &苏坦普尔SultanpurBarony{}

func init() {
	f := BSultanpur苏坦普尔.(*苏坦普尔SultanpurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sultanpur",
		TitleName: "苏坦普尔",
		TitleCode: "b_sultanpur",
	}
}
