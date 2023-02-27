package karachev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博尔霍夫BolkhovBarony struct {
	feud.BaseBarony
}

var BBolkhov博尔霍夫 feud.Barony = &博尔霍夫BolkhovBarony{}

func init() {
    f := BBolkhov博尔霍夫.(*博尔霍夫BolkhovBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bolkhov",
		TitleName: "博尔霍夫",
		TitleCode: "b_bolkhov",
	}
}
