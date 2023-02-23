package frisia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博尔斯瓦德BolswardBarony struct {
	feud.BaseBarony
}

var BBolsward博尔斯瓦德 feud.Barony = &博尔斯瓦德BolswardBarony{}

func init() {
	f := BBolsward博尔斯瓦德.(*博尔斯瓦德BolswardBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bolsward",
		TitleName: "博尔斯瓦德",
		TitleCode: "b_bolsward",
	}
}
