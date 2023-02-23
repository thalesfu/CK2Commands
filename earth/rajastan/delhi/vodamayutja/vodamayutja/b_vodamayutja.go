package vodamayutja

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 菩陀摩瑜多VodamayutjaBarony struct {
	feud.BaseBarony
}

var BVodamayutja菩陀摩瑜多 feud.Barony = &菩陀摩瑜多VodamayutjaBarony{}

func init() {
	f := BVodamayutja菩陀摩瑜多.(*菩陀摩瑜多VodamayutjaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vodamayutja",
		TitleName: "菩陀摩瑜多",
		TitleCode: "b_vodamayutja",
	}
}
