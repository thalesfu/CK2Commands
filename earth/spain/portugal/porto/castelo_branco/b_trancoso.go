package castelo_branco

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特兰科苏TrancosoBarony struct {
	feud.BaseBarony
}

var BTrancoso特兰科苏 feud.Barony = &特兰科苏TrancosoBarony{}

func init() {
    f := BTrancoso特兰科苏.(*特兰科苏TrancosoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "trancoso",
		TitleName: "特兰科苏",
		TitleCode: "b_trancoso",
	}
}
