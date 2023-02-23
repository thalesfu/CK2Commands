package ajayameru

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多罗姞利呬TaragarhBarony struct {
	feud.BaseBarony
}

var BTaragarh多罗姞利呬 feud.Barony = &多罗姞利呬TaragarhBarony{}

func init() {
	f := BTaragarh多罗姞利呬.(*多罗姞利呬TaragarhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "taragarh",
		TitleName: "多罗姞利呬",
		TitleCode: "b_taragarh",
	}
}
