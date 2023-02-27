package karor

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦巴查兰Wan_bhachranBarony struct {
	feud.BaseBarony
}

var BWan_bhachran瓦巴查兰 feud.Barony = &瓦巴查兰Wan_bhachranBarony{}

func init() {
    f := BWan_bhachran瓦巴查兰.(*瓦巴查兰Wan_bhachranBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wan_bhachran",
		TitleName: "瓦巴查兰",
		TitleCode: "b_wan_bhachran",
	}
}
