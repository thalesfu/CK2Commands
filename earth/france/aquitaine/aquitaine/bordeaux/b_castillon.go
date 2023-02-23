package bordeaux

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡斯蒂永CastillonBarony struct {
	feud.BaseBarony
}

var BCastillon卡斯蒂永 feud.Barony = &卡斯蒂永CastillonBarony{}

func init() {
	f := BCastillon卡斯蒂永.(*卡斯蒂永CastillonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "castillon",
		TitleName: "卡斯蒂永",
		TitleCode: "b_castillon",
	}
}
