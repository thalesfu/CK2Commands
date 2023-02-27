package dauphine_viennois

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦尔雷阿斯ValreasBarony struct {
	feud.BaseBarony
}

var BValreas瓦尔雷阿斯 feud.Barony = &瓦尔雷阿斯ValreasBarony{}

func init() {
    f := BValreas瓦尔雷阿斯.(*瓦尔雷阿斯ValreasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "valreas",
		TitleName: "瓦尔雷阿斯",
		TitleCode: "b_valreas",
	}
}
