package madurai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗迷湿伐蓝RamesvaramBarony struct {
	feud.BaseBarony
}

var BRamesvaram罗迷湿伐蓝 feud.Barony = &罗迷湿伐蓝RamesvaramBarony{}

func init() {
    f := BRamesvaram罗迷湿伐蓝.(*罗迷湿伐蓝RamesvaramBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ramesvaram",
		TitleName: "罗迷湿伐蓝",
		TitleCode: "b_ramesvaram",
	}
}
