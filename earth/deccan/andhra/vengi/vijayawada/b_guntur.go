package vijayawada

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贡土尔GunturBarony struct {
	feud.BaseBarony
}

var BGuntur贡土尔 feud.Barony = &贡土尔GunturBarony{}

func init() {
    f := BGuntur贡土尔.(*贡土尔GunturBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "guntur",
		TitleName: "贡土尔",
		TitleCode: "b_guntur",
	}
}
