package ludrava

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蓝耆RunjhaBarony struct {
	feud.BaseBarony
}

var BRunjha蓝耆 feud.Barony = &蓝耆RunjhaBarony{}

func init() {
	f := BRunjha蓝耆.(*蓝耆RunjhaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "runjha",
		TitleName: "蓝耆",
		TitleCode: "b_runjha",
	}
}
