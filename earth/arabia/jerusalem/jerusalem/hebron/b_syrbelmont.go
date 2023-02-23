package hebron

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝尔蒙特SyrbelmontBarony struct {
	feud.BaseBarony
}

var BSyrbelmont贝尔蒙特 feud.Barony = &贝尔蒙特SyrbelmontBarony{}

func init() {
	f := BSyrbelmont贝尔蒙特.(*贝尔蒙特SyrbelmontBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "syrbelmont",
		TitleName: "贝尔蒙特",
		TitleCode: "b_syrbelmont",
	}
}
