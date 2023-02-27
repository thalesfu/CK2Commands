package barasuru

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拘多瞿荼KothagudaBarony struct {
	feud.BaseBarony
}

var BKothaguda拘多瞿荼 feud.Barony = &拘多瞿荼KothagudaBarony{}

func init() {
    f := BKothaguda拘多瞿荼.(*拘多瞿荼KothagudaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kothaguda",
		TitleName: "拘多瞿荼",
		TitleCode: "b_kothaguda",
	}
}
