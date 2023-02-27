package nalut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杜瓦尔贝苏夫Douar_bessoufBarony struct {
	feud.BaseBarony
}

var BDouar_bessouf杜瓦尔贝苏夫 feud.Barony = &杜瓦尔贝苏夫Douar_bessoufBarony{}

func init() {
    f := BDouar_bessouf杜瓦尔贝苏夫.(*杜瓦尔贝苏夫Douar_bessoufBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "douar_bessouf",
		TitleName: "杜瓦尔贝苏夫",
		TitleCode: "b_douar_bessouf",
	}
}
