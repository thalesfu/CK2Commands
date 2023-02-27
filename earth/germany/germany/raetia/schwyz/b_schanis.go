package schwyz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 谢尼斯SchanisBarony struct {
	feud.BaseBarony
}

var BSchanis谢尼斯 feud.Barony = &谢尼斯SchanisBarony{}

func init() {
    f := BSchanis谢尼斯.(*谢尼斯SchanisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "schanis",
		TitleName: "谢尼斯",
		TitleCode: "b_schanis",
	}
}
