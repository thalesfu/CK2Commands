package qayaliq

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍尔果斯HorgosBarony struct {
	feud.BaseBarony
}

var BHorgos霍尔果斯 feud.Barony = &霍尔果斯HorgosBarony{}

func init() {
    f := BHorgos霍尔果斯.(*霍尔果斯HorgosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "horgos",
		TitleName: "霍尔果斯",
		TitleCode: "b_horgos",
	}
}
