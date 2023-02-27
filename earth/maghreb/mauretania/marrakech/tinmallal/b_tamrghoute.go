package tinmallal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泰姆尔古泰TamrghouteBarony struct {
	feud.BaseBarony
}

var BTamrghoute泰姆尔古泰 feud.Barony = &泰姆尔古泰TamrghouteBarony{}

func init() {
    f := BTamrghoute泰姆尔古泰.(*泰姆尔古泰TamrghouteBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tamrghoute",
		TitleName: "泰姆尔古泰",
		TitleCode: "b_tamrghoute",
	}
}
