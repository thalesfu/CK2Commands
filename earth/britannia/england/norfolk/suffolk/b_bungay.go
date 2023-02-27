package suffolk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 邦吉BungayBarony struct {
	feud.BaseBarony
}

var BBungay邦吉 feud.Barony = &邦吉BungayBarony{}

func init() {
    f := BBungay邦吉.(*邦吉BungayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bungay",
		TitleName: "邦吉",
		TitleCode: "b_bungay",
	}
}
