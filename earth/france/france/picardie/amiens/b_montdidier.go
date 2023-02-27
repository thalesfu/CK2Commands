package amiens

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙迪迪耶MontdidierBarony struct {
	feud.BaseBarony
}

var BMontdidier蒙迪迪耶 feud.Barony = &蒙迪迪耶MontdidierBarony{}

func init() {
    f := BMontdidier蒙迪迪耶.(*蒙迪迪耶MontdidierBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "montdidier",
		TitleName: "蒙迪迪耶",
		TitleCode: "b_montdidier",
	}
}
