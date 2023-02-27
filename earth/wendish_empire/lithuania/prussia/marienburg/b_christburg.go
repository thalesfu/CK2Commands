package marienburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基督堡ChristburgBarony struct {
	feud.BaseBarony
}

var BChristburg基督堡 feud.Barony = &基督堡ChristburgBarony{}

func init() {
    f := BChristburg基督堡.(*基督堡ChristburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "christburg",
		TitleName: "基督堡",
		TitleCode: "b_christburg",
	}
}
