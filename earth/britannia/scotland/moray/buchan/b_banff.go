package buchan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 班夫BanffBarony struct {
	feud.BaseBarony
}

var BBanff班夫 feud.Barony = &班夫BanffBarony{}

func init() {
	f := BBanff班夫.(*班夫BanffBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "banff",
		TitleName: "班夫",
		TitleCode: "b_banff",
	}
}
