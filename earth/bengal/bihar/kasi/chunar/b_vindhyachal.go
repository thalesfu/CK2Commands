package chunar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 频陀遮罗VindhyachalBarony struct {
	feud.BaseBarony
}

var BVindhyachal频陀遮罗 feud.Barony = &频陀遮罗VindhyachalBarony{}

func init() {
	f := BVindhyachal频陀遮罗.(*频陀遮罗VindhyachalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vindhyachal",
		TitleName: "频陀遮罗",
		TitleCode: "b_vindhyachal",
	}
}
