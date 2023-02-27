package neumark

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 腾佩尔堡TempelburgBarony struct {
	feud.BaseBarony
}

var BTempelburg腾佩尔堡 feud.Barony = &腾佩尔堡TempelburgBarony{}

func init() {
    f := BTempelburg腾佩尔堡.(*腾佩尔堡TempelburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tempelburg",
		TitleName: "腾佩尔堡",
		TitleCode: "b_tempelburg",
	}
}
