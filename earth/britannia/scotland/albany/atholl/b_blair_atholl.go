package atholl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布莱尔阿瑟尔Blair_athollBarony struct {
	feud.BaseBarony
}

var BBlair_atholl布莱尔阿瑟尔 feud.Barony = &布莱尔阿瑟尔Blair_athollBarony{}

func init() {
    f := BBlair_atholl布莱尔阿瑟尔.(*布莱尔阿瑟尔Blair_athollBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "blair_atholl",
		TitleName: "布莱尔阿瑟尔",
		TitleCode: "b_blair_atholl",
	}
}
