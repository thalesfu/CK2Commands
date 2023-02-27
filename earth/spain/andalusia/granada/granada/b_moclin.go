package granada

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫克林MoclinBarony struct {
	feud.BaseBarony
}

var BMoclin莫克林 feud.Barony = &莫克林MoclinBarony{}

func init() {
    f := BMoclin莫克林.(*莫克林MoclinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "moclin",
		TitleName: "莫克林",
		TitleCode: "b_moclin",
	}
}
