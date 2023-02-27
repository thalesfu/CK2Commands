package udine

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马尼亚戈ManiagoBarony struct {
	feud.BaseBarony
}

var BManiago马尼亚戈 feud.Barony = &马尼亚戈ManiagoBarony{}

func init() {
    f := BManiago马尼亚戈.(*马尼亚戈ManiagoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maniago",
		TitleName: "马尼亚戈",
		TitleCode: "b_maniago",
	}
}
