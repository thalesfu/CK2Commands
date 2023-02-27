package sakala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕萨鲁尔PasrurBarony struct {
	feud.BaseBarony
}

var BPasrur帕萨鲁尔 feud.Barony = &帕萨鲁尔PasrurBarony{}

func init() {
    f := BPasrur帕萨鲁尔.(*帕萨鲁尔PasrurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pasrur",
		TitleName: "帕萨鲁尔",
		TitleCode: "b_pasrur",
	}
}
