package srihatta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿赡补罗AzampurBarony struct {
	feud.BaseBarony
}

var BAzampur阿赡补罗 feud.Barony = &阿赡补罗AzampurBarony{}

func init() {
    f := BAzampur阿赡补罗.(*阿赡补罗AzampurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "azampur",
		TitleName: "阿赡补罗",
		TitleCode: "b_azampur",
	}
}
