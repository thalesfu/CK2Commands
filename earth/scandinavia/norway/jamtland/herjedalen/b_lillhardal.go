package herjedalen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 小海达尔LillhardalBarony struct {
	feud.BaseBarony
}

var BLillhardal小海达尔 feud.Barony = &小海达尔LillhardalBarony{}

func init() {
    f := BLillhardal小海达尔.(*小海达尔LillhardalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lillhardal",
		TitleName: "小海达尔",
		TitleCode: "b_lillhardal",
	}
}
