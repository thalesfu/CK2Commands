package lemdiyya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 提亚雷特TiaretBarony struct {
	feud.BaseBarony
}

var BTiaret提亚雷特 feud.Barony = &提亚雷特TiaretBarony{}

func init() {
	f := BTiaret提亚雷特.(*提亚雷特TiaretBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tiaret",
		TitleName: "提亚雷特",
		TitleCode: "b_tiaret",
	}
}
