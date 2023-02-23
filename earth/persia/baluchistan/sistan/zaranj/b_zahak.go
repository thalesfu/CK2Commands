package zaranj

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎哈克ZahakBarony struct {
	feud.BaseBarony
}

var BZahak扎哈克 feud.Barony = &扎哈克ZahakBarony{}

func init() {
	f := BZahak扎哈克.(*扎哈克ZahakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zahak",
		TitleName: "扎哈克",
		TitleCode: "b_zahak",
	}
}
