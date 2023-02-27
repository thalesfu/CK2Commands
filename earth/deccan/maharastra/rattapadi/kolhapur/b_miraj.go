package kolhapur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米勒杰MirajBarony struct {
	feud.BaseBarony
}

var BMiraj米勒杰 feud.Barony = &米勒杰MirajBarony{}

func init() {
    f := BMiraj米勒杰.(*米勒杰MirajBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "miraj",
		TitleName: "米勒杰",
		TitleCode: "b_miraj",
	}
}
