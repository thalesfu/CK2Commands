package ujjayini

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 邬阇衍那UjjayiniBarony struct {
	feud.BaseBarony
}

var BUjjayini邬阇衍那 feud.Barony = &邬阇衍那UjjayiniBarony{}

func init() {
    f := BUjjayini邬阇衍那.(*邬阇衍那UjjayiniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ujjayini",
		TitleName: "邬阇衍那",
		TitleCode: "b_ujjayini",
	}
}
