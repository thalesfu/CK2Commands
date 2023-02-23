package djenne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 班杜古BinduguBarony struct {
	feud.BaseBarony
}

var BBindugu班杜古 feud.Barony = &班杜古BinduguBarony{}

func init() {
	f := BBindugu班杜古.(*班杜古BinduguBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bindugu",
		TitleName: "班杜古",
		TitleCode: "b_bindugu",
	}
}
