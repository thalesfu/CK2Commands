package zarma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦拉姆OuallamBarony struct {
	feud.BaseBarony
}

var BOuallam瓦拉姆 feud.Barony = &瓦拉姆OuallamBarony{}

func init() {
    f := BOuallam瓦拉姆.(*瓦拉姆OuallamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ouallam",
		TitleName: "瓦拉姆",
		TitleCode: "b_ouallam",
	}
}
