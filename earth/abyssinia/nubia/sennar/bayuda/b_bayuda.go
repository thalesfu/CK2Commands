package bayuda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拜尤达BayudaBarony struct {
	feud.BaseBarony
}

var BBayuda拜尤达 feud.Barony = &拜尤达BayudaBarony{}

func init() {
	f := BBayuda拜尤达.(*拜尤达BayudaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bayuda",
		TitleName: "拜尤达",
		TitleCode: "b_bayuda",
	}
}
