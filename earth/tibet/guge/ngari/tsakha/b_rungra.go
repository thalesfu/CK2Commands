package tsakha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 绒热RungraBarony struct {
	feud.BaseBarony
}

var BRungra绒热 feud.Barony = &绒热RungraBarony{}

func init() {
	f := BRungra绒热.(*绒热RungraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rungra",
		TitleName: "绒热",
		TitleCode: "b_rungra",
	}
}
