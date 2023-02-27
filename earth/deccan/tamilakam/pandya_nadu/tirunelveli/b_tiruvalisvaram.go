package tirunelveli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帝卢伐利湿伐蓝TiruvalisvaramBarony struct {
	feud.BaseBarony
}

var BTiruvalisvaram帝卢伐利湿伐蓝 feud.Barony = &帝卢伐利湿伐蓝TiruvalisvaramBarony{}

func init() {
    f := BTiruvalisvaram帝卢伐利湿伐蓝.(*帝卢伐利湿伐蓝TiruvalisvaramBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tiruvalisvaram",
		TitleName: "帝卢伐利湿伐蓝",
		TitleCode: "b_tiruvalisvaram",
	}
}
