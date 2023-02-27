package luxembourg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卢森堡LuxembourgBarony struct {
	feud.BaseBarony
}

var BLuxembourg卢森堡 feud.Barony = &卢森堡LuxembourgBarony{}

func init() {
    f := BLuxembourg卢森堡.(*卢森堡LuxembourgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "luxembourg",
		TitleName: "卢森堡",
		TitleCode: "b_luxembourg",
	}
}
