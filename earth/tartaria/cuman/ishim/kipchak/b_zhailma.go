package kipchak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎伊尔马ZhailmaBarony struct {
	feud.BaseBarony
}

var BZhailma扎伊尔马 feud.Barony = &扎伊尔马ZhailmaBarony{}

func init() {
    f := BZhailma扎伊尔马.(*扎伊尔马ZhailmaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zhailma",
		TitleName: "扎伊尔马",
		TitleCode: "b_zhailma",
	}
}
