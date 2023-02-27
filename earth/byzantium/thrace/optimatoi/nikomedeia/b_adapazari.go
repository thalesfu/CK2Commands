package nikomedeia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿达帕扎里AdapazariBarony struct {
	feud.BaseBarony
}

var BAdapazari阿达帕扎里 feud.Barony = &阿达帕扎里AdapazariBarony{}

func init() {
    f := BAdapazari阿达帕扎里.(*阿达帕扎里AdapazariBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "adapazari",
		TitleName: "阿达帕扎里",
		TitleCode: "b_adapazari",
	}
}
