package djimi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉米DjimiBarony struct {
	feud.BaseBarony
}

var BDjimi吉米 feud.Barony = &吉米DjimiBarony{}

func init() {
	f := BDjimi吉米.(*吉米DjimiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "djimi",
		TitleName: "吉米",
		TitleCode: "b_djimi",
	}
}
