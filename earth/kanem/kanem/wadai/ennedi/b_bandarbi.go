package ennedi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 班达尔比BandarbiBarony struct {
	feud.BaseBarony
}

var BBandarbi班达尔比 feud.Barony = &班达尔比BandarbiBarony{}

func init() {
    f := BBandarbi班达尔比.(*班达尔比BandarbiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bandarbi",
		TitleName: "班达尔比",
		TitleCode: "b_bandarbi",
	}
}
