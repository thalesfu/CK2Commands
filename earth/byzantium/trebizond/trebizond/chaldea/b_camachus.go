package chaldea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡马库斯CamachusBarony struct {
	feud.BaseBarony
}

var BCamachus卡马库斯 feud.Barony = &卡马库斯CamachusBarony{}

func init() {
    f := BCamachus卡马库斯.(*卡马库斯CamachusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "camachus",
		TitleName: "卡马库斯",
		TitleCode: "b_camachus",
	}
}
