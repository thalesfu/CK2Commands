package galich

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 季斯梅尼齐亚TysmenytsiaBarony struct {
	feud.BaseBarony
}

var BTysmenytsia季斯梅尼齐亚 feud.Barony = &季斯梅尼齐亚TysmenytsiaBarony{}

func init() {
    f := BTysmenytsia季斯梅尼齐亚.(*季斯梅尼齐亚TysmenytsiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tysmenytsia",
		TitleName: "季斯梅尼齐亚",
		TitleCode: "b_tysmenytsia",
	}
}
