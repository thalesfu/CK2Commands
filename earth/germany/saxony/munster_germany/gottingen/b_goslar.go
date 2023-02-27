package gottingen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 戈斯拉尔GoslarBarony struct {
	feud.BaseBarony
}

var BGoslar戈斯拉尔 feud.Barony = &戈斯拉尔GoslarBarony{}

func init() {
    f := BGoslar戈斯拉尔.(*戈斯拉尔GoslarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "goslar",
		TitleName: "戈斯拉尔",
		TitleCode: "b_goslar",
	}
}
