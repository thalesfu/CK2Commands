package kucha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 张家湾ZhangjiawanBarony struct {
	feud.BaseBarony
}

var BZhangjiawan张家湾 feud.Barony = &张家湾ZhangjiawanBarony{}

func init() {
    f := BZhangjiawan张家湾.(*张家湾ZhangjiawanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zhangjiawan",
		TitleName: "张家湾",
		TitleCode: "b_zhangjiawan",
	}
}
