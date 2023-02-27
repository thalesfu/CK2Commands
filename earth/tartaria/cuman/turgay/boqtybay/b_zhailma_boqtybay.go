package boqtybay

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎伊尔马Zhailma_boqtybayBarony struct {
	feud.BaseBarony
}

var BZhailma_boqtybay扎伊尔马 feud.Barony = &扎伊尔马Zhailma_boqtybayBarony{}

func init() {
    f := BZhailma_boqtybay扎伊尔马.(*扎伊尔马Zhailma_boqtybayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zhailma_boqtybay",
		TitleName: "扎伊尔马",
		TitleCode: "b_zhailma_boqtybay",
	}
}
