package burkhan_buudai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布尔汗布岱Burkhan_buudaiBarony struct {
	feud.BaseBarony
}

var BBurkhan_buudai布尔汗布岱 feud.Barony = &布尔汗布岱Burkhan_buudaiBarony{}

func init() {
    f := BBurkhan_buudai布尔汗布岱.(*布尔汗布岱Burkhan_buudaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "burkhan_buudai",
		TitleName: "布尔汗布岱",
		TitleCode: "b_burkhan_buudai",
	}
}
