package kunggar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎雪ZaxoiBarony struct {
	feud.BaseBarony
}

var BZaxoi扎雪 feud.Barony = &扎雪ZaxoiBarony{}

func init() {
	f := BZaxoi扎雪.(*扎雪ZaxoiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zaxoi",
		TitleName: "扎雪",
		TitleCode: "b_zaxoi",
	}
}
