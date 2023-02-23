package trapani

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马尔萨拉MarsalaBarony struct {
	feud.BaseBarony
}

var BMarsala马尔萨拉 feud.Barony = &马尔萨拉MarsalaBarony{}

func init() {
	f := BMarsala马尔萨拉.(*马尔萨拉MarsalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "marsala",
		TitleName: "马尔萨拉",
		TitleCode: "b_marsala",
	}
}
