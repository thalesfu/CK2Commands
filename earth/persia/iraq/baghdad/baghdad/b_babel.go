package baghdad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴贝尔BabelBarony struct {
	feud.BaseBarony
}

var BBabel巴贝尔 feud.Barony = &巴贝尔BabelBarony{}

func init() {
	f := BBabel巴贝尔.(*巴贝尔BabelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "babel",
		TitleName: "巴贝尔",
		TitleCode: "b_babel",
	}
}
