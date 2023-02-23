package kaliopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利西马其亚LysimachiaBarony struct {
	feud.BaseBarony
}

var BLysimachia利西马其亚 feud.Barony = &利西马其亚LysimachiaBarony{}

func init() {
	f := BLysimachia利西马其亚.(*利西马其亚LysimachiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lysimachia",
		TitleName: "利西马其亚",
		TitleCode: "b_lysimachia",
	}
}
