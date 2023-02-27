package kufa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉拜什ChibayishBarony struct {
	feud.BaseBarony
}

var BChibayish吉拜什 feud.Barony = &吉拜什ChibayishBarony{}

func init() {
    f := BChibayish吉拜什.(*吉拜什ChibayishBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chibayish",
		TitleName: "吉拜什",
		TitleCode: "b_chibayish",
	}
}
