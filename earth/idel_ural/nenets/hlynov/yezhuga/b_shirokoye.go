package yezhuga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希罗科耶ShirokoyeBarony struct {
	feud.BaseBarony
}

var BShirokoye希罗科耶 feud.Barony = &希罗科耶ShirokoyeBarony{}

func init() {
    f := BShirokoye希罗科耶.(*希罗科耶ShirokoyeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shirokoye",
		TitleName: "希罗科耶",
		TitleCode: "b_shirokoye",
	}
}
