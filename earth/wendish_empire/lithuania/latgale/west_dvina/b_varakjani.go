package west_dvina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦拉克利亚尼VarakjaniBarony struct {
	feud.BaseBarony
}

var BVarakjani瓦拉克利亚尼 feud.Barony = &瓦拉克利亚尼VarakjaniBarony{}

func init() {
    f := BVarakjani瓦拉克利亚尼.(*瓦拉克利亚尼VarakjaniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "varakjani",
		TitleName: "瓦拉克利亚尼",
		TitleCode: "b_varakjani",
	}
}
