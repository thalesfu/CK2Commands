package prusa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米莱托波利斯MiletopolisBarony struct {
	feud.BaseBarony
}

var BMiletopolis米莱托波利斯 feud.Barony = &米莱托波利斯MiletopolisBarony{}

func init() {
    f := BMiletopolis米莱托波利斯.(*米莱托波利斯MiletopolisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "miletopolis",
		TitleName: "米莱托波利斯",
		TitleCode: "b_miletopolis",
	}
}
