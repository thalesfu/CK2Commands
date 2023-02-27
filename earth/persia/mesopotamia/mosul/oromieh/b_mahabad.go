package oromieh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马哈巴德MahabadBarony struct {
	feud.BaseBarony
}

var BMahabad马哈巴德 feud.Barony = &马哈巴德MahabadBarony{}

func init() {
    f := BMahabad马哈巴德.(*马哈巴德MahabadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mahabad",
		TitleName: "马哈巴德",
		TitleCode: "b_mahabad",
	}
}
