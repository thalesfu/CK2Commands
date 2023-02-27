package banavasi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吠阇演帝VaijayantiBarony struct {
	feud.BaseBarony
}

var BVaijayanti吠阇演帝 feud.Barony = &吠阇演帝VaijayantiBarony{}

func init() {
    f := BVaijayanti吠阇演帝.(*吠阇演帝VaijayantiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vaijayanti",
		TitleName: "吠阇演帝",
		TitleCode: "b_vaijayanti",
	}
}
