package kagha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 包尼BauniBarony struct {
	feud.BaseBarony
}

var BBauni包尼 feud.Barony = &包尼BauniBarony{}

func init() {
	f := BBauni包尼.(*包尼BauniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bauni",
		TitleName: "包尼",
		TitleCode: "b_bauni",
	}
}
