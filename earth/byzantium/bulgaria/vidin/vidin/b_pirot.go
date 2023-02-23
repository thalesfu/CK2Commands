package vidin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 皮罗特PirotBarony struct {
	feud.BaseBarony
}

var BPirot皮罗特 feud.Barony = &皮罗特PirotBarony{}

func init() {
	f := BPirot皮罗特.(*皮罗特PirotBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pirot",
		TitleName: "皮罗特",
		TitleCode: "b_pirot",
	}
}
