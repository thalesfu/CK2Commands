package lower_dniepr

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 别洛焦尔卡BilozerkaBarony struct {
	feud.BaseBarony
}

var BBilozerka别洛焦尔卡 feud.Barony = &别洛焦尔卡BilozerkaBarony{}

func init() {
    f := BBilozerka别洛焦尔卡.(*别洛焦尔卡BilozerkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bilozerka",
		TitleName: "别洛焦尔卡",
		TitleCode: "b_bilozerka",
	}
}
