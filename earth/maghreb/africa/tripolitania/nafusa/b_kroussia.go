package nafusa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克鲁西亚KroussiaBarony struct {
	feud.BaseBarony
}

var BKroussia克鲁西亚 feud.Barony = &克鲁西亚KroussiaBarony{}

func init() {
	f := BKroussia克鲁西亚.(*克鲁西亚KroussiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kroussia",
		TitleName: "克鲁西亚",
		TitleCode: "b_kroussia",
	}
}
