package lyubech

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯克托克SkytokBarony struct {
	feud.BaseBarony
}

var BSkytok斯克托克 feud.Barony = &斯克托克SkytokBarony{}

func init() {
	f := BSkytok斯克托克.(*斯克托克SkytokBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "skytok",
		TitleName: "斯克托克",
		TitleCode: "b_skytok",
	}
}
