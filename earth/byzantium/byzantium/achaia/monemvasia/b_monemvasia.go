package monemvasia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫奈姆瓦夏MonemvasiaBarony struct {
	feud.BaseBarony
}

var BMonemvasia莫奈姆瓦夏 feud.Barony = &莫奈姆瓦夏MonemvasiaBarony{}

func init() {
    f := BMonemvasia莫奈姆瓦夏.(*莫奈姆瓦夏MonemvasiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "monemvasia",
		TitleName: "莫奈姆瓦夏",
		TitleCode: "b_monemvasia",
	}
}
