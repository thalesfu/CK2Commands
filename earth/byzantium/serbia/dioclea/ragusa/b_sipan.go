package ragusa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希潘SipanBarony struct {
	feud.BaseBarony
}

var BSipan希潘 feud.Barony = &希潘SipanBarony{}

func init() {
	f := BSipan希潘.(*希潘SipanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sipan",
		TitleName: "希潘",
		TitleCode: "b_sipan",
	}
}
