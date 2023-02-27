package nalsia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克尔纳韦KernaveBarony struct {
	feud.BaseBarony
}

var BKernave克尔纳韦 feud.Barony = &克尔纳韦KernaveBarony{}

func init() {
    f := BKernave克尔纳韦.(*克尔纳韦KernaveBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kernave",
		TitleName: "克尔纳韦",
		TitleCode: "b_kernave",
	}
}
