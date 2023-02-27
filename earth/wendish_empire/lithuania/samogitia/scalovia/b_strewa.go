package scalovia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯特雷瓦StrewaBarony struct {
	feud.BaseBarony
}

var BStrewa斯特雷瓦 feud.Barony = &斯特雷瓦StrewaBarony{}

func init() {
    f := BStrewa斯特雷瓦.(*斯特雷瓦StrewaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "strewa",
		TitleName: "斯特雷瓦",
		TitleCode: "b_strewa",
	}
}
