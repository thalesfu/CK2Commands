package charkliq

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米兰MiranBarony struct {
	feud.BaseBarony
}

var BMiran米兰 feud.Barony = &米兰MiranBarony{}

func init() {
    f := BMiran米兰.(*米兰MiranBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "miran",
		TitleName: "米兰",
		TitleCode: "b_miran",
	}
}
