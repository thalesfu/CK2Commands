package vukovar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦尔卡ValkaBarony struct {
	feud.BaseBarony
}

var BValka瓦尔卡 feud.Barony = &瓦尔卡ValkaBarony{}

func init() {
	f := BValka瓦尔卡.(*瓦尔卡ValkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "valka",
		TitleName: "瓦尔卡",
		TitleCode: "b_valka",
	}
}
