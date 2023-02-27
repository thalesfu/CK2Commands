package orangallu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波罗姆佩特PalampetBarony struct {
	feud.BaseBarony
}

var BPalampet波罗姆佩特 feud.Barony = &波罗姆佩特PalampetBarony{}

func init() {
    f := BPalampet波罗姆佩特.(*波罗姆佩特PalampetBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "palampet",
		TitleName: "波罗姆佩特",
		TitleCode: "b_palampet",
	}
}
