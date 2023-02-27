package grassland_cheremisa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 切尔姆ChermBarony struct {
	feud.BaseBarony
}

var BCherm切尔姆 feud.Barony = &切尔姆ChermBarony{}

func init() {
    f := BCherm切尔姆.(*切尔姆ChermBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cherm",
		TitleName: "切尔姆",
		TitleCode: "b_cherm",
	}
}
