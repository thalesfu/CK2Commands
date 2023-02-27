package syria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎巴JarbaBarony struct {
	feud.BaseBarony
}

var BJarba扎巴 feud.Barony = &扎巴JarbaBarony{}

func init() {
    f := BJarba扎巴.(*扎巴JarbaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jarba",
		TitleName: "扎巴",
		TitleCode: "b_jarba",
	}
}
