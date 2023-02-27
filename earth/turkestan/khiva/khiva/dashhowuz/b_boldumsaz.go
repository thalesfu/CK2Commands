package dashhowuz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博尔杜姆萨兹BoldumsazBarony struct {
	feud.BaseBarony
}

var BBoldumsaz博尔杜姆萨兹 feud.Barony = &博尔杜姆萨兹BoldumsazBarony{}

func init() {
    f := BBoldumsaz博尔杜姆萨兹.(*博尔杜姆萨兹BoldumsazBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "boldumsaz",
		TitleName: "博尔杜姆萨兹",
		TitleCode: "b_boldumsaz",
	}
}
