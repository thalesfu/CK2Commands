package leventina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙特贝洛Montebello_leventinaBarony struct {
	feud.BaseBarony
}

var BMontebello_leventina蒙特贝洛 feud.Barony = &蒙特贝洛Montebello_leventinaBarony{}

func init() {
    f := BMontebello_leventina蒙特贝洛.(*蒙特贝洛Montebello_leventinaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "montebello_leventina",
		TitleName: "蒙特贝洛",
		TitleCode: "b_montebello_leventina",
	}
}
