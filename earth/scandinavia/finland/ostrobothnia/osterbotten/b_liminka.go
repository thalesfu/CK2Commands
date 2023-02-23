package osterbotten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利明卡LiminkaBarony struct {
	feud.BaseBarony
}

var BLiminka利明卡 feud.Barony = &利明卡LiminkaBarony{}

func init() {
	f := BLiminka利明卡.(*利明卡LiminkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "liminka",
		TitleName: "利明卡",
		TitleCode: "b_liminka",
	}
}
