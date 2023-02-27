package derby

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博尔索弗BolsoverBarony struct {
	feud.BaseBarony
}

var BBolsover博尔索弗 feud.Barony = &博尔索弗BolsoverBarony{}

func init() {
    f := BBolsover博尔索弗.(*博尔索弗BolsoverBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bolsover",
		TitleName: "博尔索弗",
		TitleCode: "b_bolsover",
	}
}
