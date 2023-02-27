package zhongba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉拉Gela_bBarony struct {
	feud.BaseBarony
}

var BGela_b吉拉 feud.Barony = &吉拉Gela_bBarony{}

func init() {
    f := BGela_b吉拉.(*吉拉Gela_bBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gela_b",
		TitleName: "吉拉",
		TitleCode: "b_gela_b",
	}
}
