package lower_don

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯科罗霍德SkorokhodBarony struct {
	feud.BaseBarony
}

var BSkorokhod斯科罗霍德 feud.Barony = &斯科罗霍德SkorokhodBarony{}

func init() {
    f := BSkorokhod斯科罗霍德.(*斯科罗霍德SkorokhodBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "skorokhod",
		TitleName: "斯科罗霍德",
		TitleCode: "b_skorokhod",
	}
}
