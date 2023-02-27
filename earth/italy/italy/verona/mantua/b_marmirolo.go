package mantua

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马尔米罗洛MarmiroloBarony struct {
	feud.BaseBarony
}

var BMarmirolo马尔米罗洛 feud.Barony = &马尔米罗洛MarmiroloBarony{}

func init() {
    f := BMarmirolo马尔米罗洛.(*马尔米罗洛MarmiroloBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "marmirolo",
		TitleName: "马尔米罗洛",
		TitleCode: "b_marmirolo",
	}
}
