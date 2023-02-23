package como

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦雷泽VareseBarony struct {
	feud.BaseBarony
}

var BVarese瓦雷泽 feud.Barony = &瓦雷泽VareseBarony{}

func init() {
	f := BVarese瓦雷泽.(*瓦雷泽VareseBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "varese",
		TitleName: "瓦雷泽",
		TitleCode: "b_varese",
	}
}
