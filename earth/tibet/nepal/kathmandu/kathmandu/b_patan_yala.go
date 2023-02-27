package kathmandu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕坦亚拉Patan_yalaBarony struct {
	feud.BaseBarony
}

var BPatan_yala帕坦亚拉 feud.Barony = &帕坦亚拉Patan_yalaBarony{}

func init() {
    f := BPatan_yala帕坦亚拉.(*帕坦亚拉Patan_yalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "patan_yala",
		TitleName: "帕坦亚拉",
		TitleCode: "b_patan_yala",
	}
}
