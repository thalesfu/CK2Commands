package mughalzhar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博尔加尔卡BolgarkaBarony struct {
	feud.BaseBarony
}

var BBolgarka博尔加尔卡 feud.Barony = &博尔加尔卡BolgarkaBarony{}

func init() {
    f := BBolgarka博尔加尔卡.(*博尔加尔卡BolgarkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bolgarka",
		TitleName: "博尔加尔卡",
		TitleCode: "b_bolgarka",
	}
}
