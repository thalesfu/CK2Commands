package tyrus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙福尔SyrmontfortBarony struct {
	feud.BaseBarony
}

var BSyrmontfort蒙福尔 feud.Barony = &蒙福尔SyrmontfortBarony{}

func init() {
	f := BSyrmontfort蒙福尔.(*蒙福尔SyrmontfortBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "syrmontfort",
		TitleName: "蒙福尔",
		TitleCode: "b_syrmontfort",
	}
}
