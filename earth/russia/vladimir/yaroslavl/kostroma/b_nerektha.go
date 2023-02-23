package kostroma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 涅列赫塔NerekthaBarony struct {
	feud.BaseBarony
}

var BNerektha涅列赫塔 feud.Barony = &涅列赫塔NerekthaBarony{}

func init() {
	f := BNerektha涅列赫塔.(*涅列赫塔NerekthaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nerektha",
		TitleName: "涅列赫塔",
		TitleCode: "b_nerektha",
	}
}
