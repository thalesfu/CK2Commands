package west_dvina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卢扎LudzaBarony struct {
	feud.BaseBarony
}

var BLudza卢扎 feud.Barony = &卢扎LudzaBarony{}

func init() {
    f := BLudza卢扎.(*卢扎LudzaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ludza",
		TitleName: "卢扎",
		TitleCode: "b_ludza",
	}
}
