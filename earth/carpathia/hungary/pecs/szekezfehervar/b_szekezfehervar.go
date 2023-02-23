package szekezfehervar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞克什白堡SzekezfehervarBarony struct {
	feud.BaseBarony
}

var BSzekezfehervar塞克什白堡 feud.Barony = &塞克什白堡SzekezfehervarBarony{}

func init() {
	f := BSzekezfehervar塞克什白堡.(*塞克什白堡SzekezfehervarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "szekezfehervar",
		TitleName: "塞克什白堡",
		TitleCode: "b_szekezfehervar",
	}
}
