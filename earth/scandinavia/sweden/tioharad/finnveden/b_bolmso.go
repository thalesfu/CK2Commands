package finnveden

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博尔姆瑟BolmsoBarony struct {
	feud.BaseBarony
}

var BBolmso博尔姆瑟 feud.Barony = &博尔姆瑟BolmsoBarony{}

func init() {
	f := BBolmso博尔姆瑟.(*博尔姆瑟BolmsoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bolmso",
		TitleName: "博尔姆瑟",
		TitleCode: "b_bolmso",
	}
}
