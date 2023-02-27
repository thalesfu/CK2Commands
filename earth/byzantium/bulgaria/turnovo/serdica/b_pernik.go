package serdica

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩尔尼克PernikBarony struct {
	feud.BaseBarony
}

var BPernik佩尔尼克 feud.Barony = &佩尔尼克PernikBarony{}

func init() {
    f := BPernik佩尔尼克.(*佩尔尼克PernikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pernik",
		TitleName: "佩尔尼克",
		TitleCode: "b_pernik",
	}
}
