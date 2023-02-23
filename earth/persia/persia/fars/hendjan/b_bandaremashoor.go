package hendjan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 班达尔马索尔BandaremashoorBarony struct {
	feud.BaseBarony
}

var BBandaremashoor班达尔马索尔 feud.Barony = &班达尔马索尔BandaremashoorBarony{}

func init() {
	f := BBandaremashoor班达尔马索尔.(*班达尔马索尔BandaremashoorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bandaremashoor",
		TitleName: "班达尔马索尔",
		TitleCode: "b_bandaremashoor",
	}
}
