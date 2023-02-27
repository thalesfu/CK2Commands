package mantua

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴尼奥洛圣维托BagnolosanvitoBarony struct {
	feud.BaseBarony
}

var BBagnolosanvito巴尼奥洛圣维托 feud.Barony = &巴尼奥洛圣维托BagnolosanvitoBarony{}

func init() {
    f := BBagnolosanvito巴尼奥洛圣维托.(*巴尼奥洛圣维托BagnolosanvitoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bagnolosanvito",
		TitleName: "巴尼奥洛圣维托",
		TitleCode: "b_bagnolosanvito",
	}
}
