package borovichi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃多贡VodogonBarony struct {
	feud.BaseBarony
}

var BVodogon沃多贡 feud.Barony = &沃多贡VodogonBarony{}

func init() {
    f := BVodogon沃多贡.(*沃多贡VodogonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vodogon",
		TitleName: "沃多贡",
		TitleCode: "b_vodogon",
	}
}
