package arigh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃吉卢AjluBarony struct {
	feud.BaseBarony
}

var BAjlu埃吉卢 feud.Barony = &埃吉卢AjluBarony{}

func init() {
	f := BAjlu埃吉卢.(*埃吉卢AjluBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ajlu",
		TitleName: "埃吉卢",
		TitleCode: "b_ajlu",
	}
}
