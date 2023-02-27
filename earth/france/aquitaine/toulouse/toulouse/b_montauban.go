package toulouse

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙托邦MontaubanBarony struct {
	feud.BaseBarony
}

var BMontauban蒙托邦 feud.Barony = &蒙托邦MontaubanBarony{}

func init() {
    f := BMontauban蒙托邦.(*蒙托邦MontaubanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "montauban",
		TitleName: "蒙托邦",
		TitleCode: "b_montauban",
	}
}
