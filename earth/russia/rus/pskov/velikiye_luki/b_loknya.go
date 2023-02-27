package velikiye_luki

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 洛克尼亚LoknyaBarony struct {
	feud.BaseBarony
}

var BLoknya洛克尼亚 feud.Barony = &洛克尼亚LoknyaBarony{}

func init() {
    f := BLoknya洛克尼亚.(*洛克尼亚LoknyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "loknya",
		TitleName: "洛克尼亚",
		TitleCode: "b_loknya",
	}
}
