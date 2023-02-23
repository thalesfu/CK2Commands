package armagnac

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利勒LisleBarony struct {
	feud.BaseBarony
}

var BLisle利勒 feud.Barony = &利勒LisleBarony{}

func init() {
	f := BLisle利勒.(*利勒LisleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lisle",
		TitleName: "利勒",
		TitleCode: "b_lisle",
	}
}
