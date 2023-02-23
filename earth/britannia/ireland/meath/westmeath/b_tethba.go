package westmeath

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泰斯韦TethbaBarony struct {
	feud.BaseBarony
}

var BTethba泰斯韦 feud.Barony = &泰斯韦TethbaBarony{}

func init() {
	f := BTethba泰斯韦.(*泰斯韦TethbaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tethba",
		TitleName: "泰斯韦",
		TitleCode: "b_tethba",
	}
}
