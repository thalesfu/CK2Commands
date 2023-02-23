package kumul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 辟展PiqanBarony struct {
	feud.BaseBarony
}

var BPiqan辟展 feud.Barony = &辟展PiqanBarony{}

func init() {
	f := BPiqan辟展.(*辟展PiqanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "piqan",
		TitleName: "辟展",
		TitleCode: "b_piqan",
	}
}
