package nubia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拿法布NafabBarony struct {
	feud.BaseBarony
}

var BNafab拿法布 feud.Barony = &拿法布NafabBarony{}

func init() {
	f := BNafab拿法布.(*拿法布NafabBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nafab",
		TitleName: "拿法布",
		TitleCode: "b_nafab",
	}
}
