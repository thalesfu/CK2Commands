package niederbayern

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维特尔斯巴赫WittelsbachBarony struct {
	feud.BaseBarony
}

var BWittelsbach维特尔斯巴赫 feud.Barony = &维特尔斯巴赫WittelsbachBarony{}

func init() {
	f := BWittelsbach维特尔斯巴赫.(*维特尔斯巴赫WittelsbachBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wittelsbach",
		TitleName: "维特尔斯巴赫",
		TitleCode: "b_wittelsbach",
	}
}
