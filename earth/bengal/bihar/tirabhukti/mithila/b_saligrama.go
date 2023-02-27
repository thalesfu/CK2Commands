package mithila

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 舍利伽罗摩SaligramaBarony struct {
	feud.BaseBarony
}

var BSaligrama舍利伽罗摩 feud.Barony = &舍利伽罗摩SaligramaBarony{}

func init() {
    f := BSaligrama舍利伽罗摩.(*舍利伽罗摩SaligramaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saligrama",
		TitleName: "舍利伽罗摩",
		TitleCode: "b_saligrama",
	}
}
