package soria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科瓦莱达CovaledaBarony struct {
	feud.BaseBarony
}

var BCovaleda科瓦莱达 feud.Barony = &科瓦莱达CovaledaBarony{}

func init() {
    f := BCovaleda科瓦莱达.(*科瓦莱达CovaledaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "covaleda",
		TitleName: "科瓦莱达",
		TitleCode: "b_covaleda",
	}
}
