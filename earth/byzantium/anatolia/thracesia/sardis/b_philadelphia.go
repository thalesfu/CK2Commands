package sardis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 菲拉德尔菲亚PhiladelphiaBarony struct {
	feud.BaseBarony
}

var BPhiladelphia菲拉德尔菲亚 feud.Barony = &菲拉德尔菲亚PhiladelphiaBarony{}

func init() {
    f := BPhiladelphia菲拉德尔菲亚.(*菲拉德尔菲亚PhiladelphiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "philadelphia",
		TitleName: "菲拉德尔菲亚",
		TitleCode: "b_philadelphia",
	}
}
