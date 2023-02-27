package olvia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥利维亚OlviaBarony struct {
	feud.BaseBarony
}

var BOlvia奥利维亚 feud.Barony = &奥利维亚OlviaBarony{}

func init() {
    f := BOlvia奥利维亚.(*奥利维亚OlviaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "olvia",
		TitleName: "奥利维亚",
		TitleCode: "b_olvia",
	}
}
