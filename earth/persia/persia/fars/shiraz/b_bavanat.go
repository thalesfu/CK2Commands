package shiraz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴瓦纳特BavanatBarony struct {
	feud.BaseBarony
}

var BBavanat巴瓦纳特 feud.Barony = &巴瓦纳特BavanatBarony{}

func init() {
	f := BBavanat巴瓦纳特.(*巴瓦纳特BavanatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bavanat",
		TitleName: "巴瓦纳特",
		TitleCode: "b_bavanat",
	}
}
