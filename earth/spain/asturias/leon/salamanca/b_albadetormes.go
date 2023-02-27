package salamanca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托梅斯河畔阿尔瓦AlbadetormesBarony struct {
	feud.BaseBarony
}

var BAlbadetormes托梅斯河畔阿尔瓦 feud.Barony = &托梅斯河畔阿尔瓦AlbadetormesBarony{}

func init() {
    f := BAlbadetormes托梅斯河畔阿尔瓦.(*托梅斯河畔阿尔瓦AlbadetormesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "albadetormes",
		TitleName: "托梅斯河畔阿尔瓦",
		TitleCode: "b_albadetormes",
	}
}
