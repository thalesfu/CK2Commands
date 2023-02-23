package tyumen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诺夫塔普NovtapBarony struct {
	feud.BaseBarony
}

var BNovtap诺夫塔普 feud.Barony = &诺夫塔普NovtapBarony{}

func init() {
	f := BNovtap诺夫塔普.(*诺夫塔普NovtapBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "novtap",
		TitleName: "诺夫塔普",
		TitleCode: "b_novtap",
	}
}
