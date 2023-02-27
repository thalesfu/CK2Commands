package zavarot

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎瓦罗特ZavarotBarony struct {
	feud.BaseBarony
}

var BZavarot扎瓦罗特 feud.Barony = &扎瓦罗特ZavarotBarony{}

func init() {
    f := BZavarot扎瓦罗特.(*扎瓦罗特ZavarotBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zavarot",
		TitleName: "扎瓦罗特",
		TitleCode: "b_zavarot",
	}
}
