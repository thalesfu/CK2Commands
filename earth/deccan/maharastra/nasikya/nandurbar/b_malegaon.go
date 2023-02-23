package nandurbar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩戾伽罗摩MalegaonBarony struct {
	feud.BaseBarony
}

var BMalegaon摩戾伽罗摩 feud.Barony = &摩戾伽罗摩MalegaonBarony{}

func init() {
	f := BMalegaon摩戾伽罗摩.(*摩戾伽罗摩MalegaonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "malegaon",
		TitleName: "摩戾伽罗摩",
		TitleCode: "b_malegaon",
	}
}
