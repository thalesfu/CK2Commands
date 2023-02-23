package ob

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃洛达尔卡VolodarkaBarony struct {
	feud.BaseBarony
}

var BVolodarka沃洛达尔卡 feud.Barony = &沃洛达尔卡VolodarkaBarony{}

func init() {
	f := BVolodarka沃洛达尔卡.(*沃洛达尔卡VolodarkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "volodarka",
		TitleName: "沃洛达尔卡",
		TitleCode: "b_volodarka",
	}
}
