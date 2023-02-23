package muztau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌伦海卡UrunkhaykaBarony struct {
	feud.BaseBarony
}

var BUrunkhayka乌伦海卡 feud.Barony = &乌伦海卡UrunkhaykaBarony{}

func init() {
	f := BUrunkhayka乌伦海卡.(*乌伦海卡UrunkhaykaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "urunkhayka",
		TitleName: "乌伦海卡",
		TitleCode: "b_urunkhayka",
	}
}
