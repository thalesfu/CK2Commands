package velikiye_luki

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卢基VelikiyelukiBarony struct {
	feud.BaseBarony
}

var BVelikiyeluki卢基 feud.Barony = &卢基VelikiyelukiBarony{}

func init() {
    f := BVelikiyeluki卢基.(*卢基VelikiyelukiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "velikiyeluki",
		TitleName: "卢基",
		TitleCode: "b_velikiyeluki",
	}
}
