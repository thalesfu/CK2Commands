package ugra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃德内VodnyyBarony struct {
	feud.BaseBarony
}

var BVodnyy沃德内 feud.Barony = &沃德内VodnyyBarony{}

func init() {
    f := BVodnyy沃德内.(*沃德内VodnyyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vodnyy",
		TitleName: "沃德内",
		TitleCode: "b_vodnyy",
	}
}
