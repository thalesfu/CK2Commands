package yungguan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 陀头TuotouBarony struct {
	feud.BaseBarony
}

var BTuotou陀头 feud.Barony = &陀头TuotouBarony{}

func init() {
	f := BTuotou陀头.(*陀头TuotouBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tuotou",
		TitleName: "陀头",
		TitleCode: "b_tuotou",
	}
}
