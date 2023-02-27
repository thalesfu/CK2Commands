package szekezfehervar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 毛尔曹利MarcaliBarony struct {
	feud.BaseBarony
}

var BMarcali毛尔曹利 feud.Barony = &毛尔曹利MarcaliBarony{}

func init() {
    f := BMarcali毛尔曹利.(*毛尔曹利MarcaliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "marcali",
		TitleName: "毛尔曹利",
		TitleCode: "b_marcali",
	}
}
