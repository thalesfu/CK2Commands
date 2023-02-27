package minsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马里纳戈尔卡Maryina_horkaBarony struct {
	feud.BaseBarony
}

var BMaryina_horka马里纳戈尔卡 feud.Barony = &马里纳戈尔卡Maryina_horkaBarony{}

func init() {
    f := BMaryina_horka马里纳戈尔卡.(*马里纳戈尔卡Maryina_horkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maryina_horka",
		TitleName: "马里纳戈尔卡",
		TitleCode: "b_maryina_horka",
	}
}
