package baden

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杜尔拉赫DurlachBarony struct {
	feud.BaseBarony
}

var BDurlach杜尔拉赫 feud.Barony = &杜尔拉赫DurlachBarony{}

func init() {
    f := BDurlach杜尔拉赫.(*杜尔拉赫DurlachBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "durlach",
		TitleName: "杜尔拉赫",
		TitleCode: "b_durlach",
	}
}
