package zagreb

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯泰尼奇尼亚克StenicnjakBarony struct {
	feud.BaseBarony
}

var BStenicnjak斯泰尼奇尼亚克 feud.Barony = &斯泰尼奇尼亚克StenicnjakBarony{}

func init() {
    f := BStenicnjak斯泰尼奇尼亚克.(*斯泰尼奇尼亚克StenicnjakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stenicnjak",
		TitleName: "斯泰尼奇尼亚克",
		TitleCode: "b_stenicnjak",
	}
}
