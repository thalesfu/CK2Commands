package kartli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博尔尼西BolnisiBarony struct {
	feud.BaseBarony
}

var BBolnisi博尔尼西 feud.Barony = &博尔尼西BolnisiBarony{}

func init() {
	f := BBolnisi博尔尼西.(*博尔尼西BolnisiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bolnisi",
		TitleName: "博尔尼西",
		TitleCode: "b_bolnisi",
	}
}
