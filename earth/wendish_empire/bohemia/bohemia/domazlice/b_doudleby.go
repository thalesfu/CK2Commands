package domazlice

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多乌德莱比DoudlebyBarony struct {
	feud.BaseBarony
}

var BDoudleby多乌德莱比 feud.Barony = &多乌德莱比DoudlebyBarony{}

func init() {
    f := BDoudleby多乌德莱比.(*多乌德莱比DoudlebyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "doudleby",
		TitleName: "多乌德莱比",
		TitleCode: "b_doudleby",
	}
}
