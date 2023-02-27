package medjybij

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 柳巴尔LiubarBarony struct {
	feud.BaseBarony
}

var BLiubar柳巴尔 feud.Barony = &柳巴尔LiubarBarony{}

func init() {
    f := BLiubar柳巴尔.(*柳巴尔LiubarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "liubar",
		TitleName: "柳巴尔",
		TitleCode: "b_liubar",
	}
}
