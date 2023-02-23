package gwent

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡利恩CaerleonBarony struct {
	feud.BaseBarony
}

var BCaerleon卡利恩 feud.Barony = &卡利恩CaerleonBarony{}

func init() {
	f := BCaerleon卡利恩.(*卡利恩CaerleonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "caerleon",
		TitleName: "卡利恩",
		TitleCode: "b_caerleon",
	}
}
