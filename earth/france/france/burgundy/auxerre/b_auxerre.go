package auxerre

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 欧塞尔AuxerreBarony struct {
	feud.BaseBarony
}

var BAuxerre欧塞尔 feud.Barony = &欧塞尔AuxerreBarony{}

func init() {
	f := BAuxerre欧塞尔.(*欧塞尔AuxerreBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "auxerre",
		TitleName: "欧塞尔",
		TitleCode: "b_auxerre",
	}
}
