package ifni

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔尔法亚TartayaBarony struct {
	feud.BaseBarony
}

var BTartaya塔尔法亚 feud.Barony = &塔尔法亚TartayaBarony{}

func init() {
	f := BTartaya塔尔法亚.(*塔尔法亚TartayaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tartaya",
		TitleName: "塔尔法亚",
		TitleCode: "b_tartaya",
	}
}
