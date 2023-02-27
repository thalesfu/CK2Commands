package la_mancha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍尔克拉JorqueraBarony struct {
	feud.BaseBarony
}

var BJorquera霍尔克拉 feud.Barony = &霍尔克拉JorqueraBarony{}

func init() {
    f := BJorquera霍尔克拉.(*霍尔克拉JorqueraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jorquera",
		TitleName: "霍尔克拉",
		TitleCode: "b_jorquera",
	}
}
