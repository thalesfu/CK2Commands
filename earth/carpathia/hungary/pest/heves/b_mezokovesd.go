package heves

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迈泽克韦什德MezokovesdBarony struct {
	feud.BaseBarony
}

var BMezokovesd迈泽克韦什德 feud.Barony = &迈泽克韦什德MezokovesdBarony{}

func init() {
	f := BMezokovesd迈泽克韦什德.(*迈泽克韦什德MezokovesdBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mezokovesd",
		TitleName: "迈泽克韦什德",
		TitleCode: "b_mezokovesd",
	}
}
