package feher

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 久洛白堡GyulafehervarBarony struct {
	feud.BaseBarony
}

var BGyulafehervar久洛白堡 feud.Barony = &久洛白堡GyulafehervarBarony{}

func init() {
	f := BGyulafehervar久洛白堡.(*久洛白堡GyulafehervarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gyulafehervar",
		TitleName: "久洛白堡",
		TitleCode: "b_gyulafehervar",
	}
}
