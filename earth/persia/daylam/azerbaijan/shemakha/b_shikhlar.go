package shemakha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 什克拉尔ShikhlarBarony struct {
	feud.BaseBarony
}

var BShikhlar什克拉尔 feud.Barony = &什克拉尔ShikhlarBarony{}

func init() {
    f := BShikhlar什克拉尔.(*什克拉尔ShikhlarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shikhlar",
		TitleName: "什克拉尔",
		TitleCode: "b_shikhlar",
	}
}
