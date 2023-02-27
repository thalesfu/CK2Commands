package kajaani

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基耶希梅KiehimanBarony struct {
	feud.BaseBarony
}

var BKiehiman基耶希梅 feud.Barony = &基耶希梅KiehimanBarony{}

func init() {
    f := BKiehiman基耶希梅.(*基耶希梅KiehimanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kiehiman",
		TitleName: "基耶希梅",
		TitleCode: "b_kiehiman",
	}
}
