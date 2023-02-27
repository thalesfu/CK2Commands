package khaybar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉比格RabigBarony struct {
	feud.BaseBarony
}

var BRabig拉比格 feud.Barony = &拉比格RabigBarony{}

func init() {
    f := BRabig拉比格.(*拉比格RabigBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rabig",
		TitleName: "拉比格",
		TitleCode: "b_rabig",
	}
}
