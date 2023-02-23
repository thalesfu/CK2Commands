package otgontenger

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古尔班布拉格GurvanbulagBarony struct {
	feud.BaseBarony
}

var BGurvanbulag古尔班布拉格 feud.Barony = &古尔班布拉格GurvanbulagBarony{}

func init() {
	f := BGurvanbulag古尔班布拉格.(*古尔班布拉格GurvanbulagBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gurvanbulag",
		TitleName: "古尔班布拉格",
		TitleCode: "b_gurvanbulag",
	}
}
