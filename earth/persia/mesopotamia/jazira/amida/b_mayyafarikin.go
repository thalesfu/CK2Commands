package amida

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迈亚法里金MayyafarikinBarony struct {
	feud.BaseBarony
}

var BMayyafarikin迈亚法里金 feud.Barony = &迈亚法里金MayyafarikinBarony{}

func init() {
	f := BMayyafarikin迈亚法里金.(*迈亚法里金MayyafarikinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mayyafarikin",
		TitleName: "迈亚法里金",
		TitleCode: "b_mayyafarikin",
	}
}
