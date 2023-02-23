package annaba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯基克达SkikdaBarony struct {
	feud.BaseBarony
}

var BSkikda斯基克达 feud.Barony = &斯基克达SkikdaBarony{}

func init() {
	f := BSkikda斯基克达.(*斯基克达SkikdaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "skikda",
		TitleName: "斯基克达",
		TitleCode: "b_skikda",
	}
}
