package vista

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 内斯堡NasborgBarony struct {
	feud.BaseBarony
}

var BNasborg内斯堡 feud.Barony = &内斯堡NasborgBarony{}

func init() {
    f := BNasborg内斯堡.(*内斯堡NasborgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nasborg",
		TitleName: "内斯堡",
		TitleCode: "b_nasborg",
	}
}
