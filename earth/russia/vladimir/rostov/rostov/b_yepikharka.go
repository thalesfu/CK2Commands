package rostov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 叶皮哈尔卡YepikharkaBarony struct {
	feud.BaseBarony
}

var BYepikharka叶皮哈尔卡 feud.Barony = &叶皮哈尔卡YepikharkaBarony{}

func init() {
	f := BYepikharka叶皮哈尔卡.(*叶皮哈尔卡YepikharkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yepikharka",
		TitleName: "叶皮哈尔卡",
		TitleCode: "b_yepikharka",
	}
}
