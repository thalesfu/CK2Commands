package middlesex

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托特纳姆TottenhamBarony struct {
	feud.BaseBarony
}

var BTottenham托特纳姆 feud.Barony = &托特纳姆TottenhamBarony{}

func init() {
    f := BTottenham托特纳姆.(*托特纳姆TottenhamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tottenham",
		TitleName: "托特纳姆",
		TitleCode: "b_tottenham",
	}
}
