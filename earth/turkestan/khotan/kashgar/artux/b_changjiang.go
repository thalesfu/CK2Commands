package artux

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 昌疆ChangjiangBarony struct {
	feud.BaseBarony
}

var BChangjiang昌疆 feud.Barony = &昌疆ChangjiangBarony{}

func init() {
	f := BChangjiang昌疆.(*昌疆ChangjiangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "changjiang",
		TitleName: "昌疆",
		TitleCode: "b_changjiang",
	}
}
