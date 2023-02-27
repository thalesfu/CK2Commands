package sarqihya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布巴斯提斯BubastisBarony struct {
	feud.BaseBarony
}

var BBubastis布巴斯提斯 feud.Barony = &布巴斯提斯BubastisBarony{}

func init() {
    f := BBubastis布巴斯提斯.(*布巴斯提斯BubastisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bubastis",
		TitleName: "布巴斯提斯",
		TitleCode: "b_bubastis",
	}
}
