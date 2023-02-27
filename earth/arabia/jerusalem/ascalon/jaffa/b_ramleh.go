package jaffa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉姆拉RamlehBarony struct {
	feud.BaseBarony
}

var BRamleh拉姆拉 feud.Barony = &拉姆拉RamlehBarony{}

func init() {
    f := BRamleh拉姆拉.(*拉姆拉RamlehBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ramleh",
		TitleName: "拉姆拉",
		TitleCode: "b_ramleh",
	}
}
