package sakya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 陈塘ZhentangBarony struct {
	feud.BaseBarony
}

var BZhentang陈塘 feud.Barony = &陈塘ZhentangBarony{}

func init() {
	f := BZhentang陈塘.(*陈塘ZhentangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zhentang",
		TitleName: "陈塘",
		TitleCode: "b_zhentang",
	}
}
