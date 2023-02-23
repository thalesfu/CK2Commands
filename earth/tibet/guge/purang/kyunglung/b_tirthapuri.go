package kyunglung

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 直达布日TirthapuriBarony struct {
	feud.BaseBarony
}

var BTirthapuri直达布日 feud.Barony = &直达布日TirthapuriBarony{}

func init() {
	f := BTirthapuri直达布日.(*直达布日TirthapuriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tirthapuri",
		TitleName: "直达布日",
		TitleCode: "b_tirthapuri",
	}
}
