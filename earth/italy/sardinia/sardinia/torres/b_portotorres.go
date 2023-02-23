package torres

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托雷斯港PortotorresBarony struct {
	feud.BaseBarony
}

var BPortotorres托雷斯港 feud.Barony = &托雷斯港PortotorresBarony{}

func init() {
	f := BPortotorres托雷斯港.(*托雷斯港PortotorresBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "portotorres",
		TitleName: "托雷斯港",
		TitleCode: "b_portotorres",
	}
}
