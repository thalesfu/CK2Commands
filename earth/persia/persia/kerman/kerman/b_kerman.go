package kerman

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克尔曼KermanBarony struct {
	feud.BaseBarony
}

var BKerman克尔曼 feud.Barony = &克尔曼KermanBarony{}

func init() {
	f := BKerman克尔曼.(*克尔曼KermanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kerman",
		TitleName: "克尔曼",
		TitleCode: "b_kerman",
	}
}
