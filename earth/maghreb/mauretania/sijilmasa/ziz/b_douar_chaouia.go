package ziz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杜瓦尔沙威亚Douar_chaouiaBarony struct {
	feud.BaseBarony
}

var BDouar_chaouia杜瓦尔沙威亚 feud.Barony = &杜瓦尔沙威亚Douar_chaouiaBarony{}

func init() {
    f := BDouar_chaouia杜瓦尔沙威亚.(*杜瓦尔沙威亚Douar_chaouiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "douar_chaouia",
		TitleName: "杜瓦尔沙威亚",
		TitleCode: "b_douar_chaouia",
	}
}
