package ascalon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉佛比LaforbieBarony struct {
	feud.BaseBarony
}

var BLaforbie拉佛比 feud.Barony = &拉佛比LaforbieBarony{}

func init() {
	f := BLaforbie拉佛比.(*拉佛比LaforbieBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "laforbie",
		TitleName: "拉佛比",
		TitleCode: "b_laforbie",
	}
}
