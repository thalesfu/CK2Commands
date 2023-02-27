package bamiyan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎哈克ZakhakBarony struct {
	feud.BaseBarony
}

var BZakhak扎哈克 feud.Barony = &扎哈克ZakhakBarony{}

func init() {
    f := BZakhak扎哈克.(*扎哈克ZakhakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zakhak",
		TitleName: "扎哈克",
		TitleCode: "b_zakhak",
	}
}
