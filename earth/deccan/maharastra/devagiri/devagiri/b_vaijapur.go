package devagiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吠阇补罗VaijapurBarony struct {
	feud.BaseBarony
}

var BVaijapur吠阇补罗 feud.Barony = &吠阇补罗VaijapurBarony{}

func init() {
	f := BVaijapur吠阇补罗.(*吠阇补罗VaijapurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vaijapur",
		TitleName: "吠阇补罗",
		TitleCode: "b_vaijapur",
	}
}
