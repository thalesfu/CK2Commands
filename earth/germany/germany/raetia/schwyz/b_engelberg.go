package schwyz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 恩格尔贝格EngelbergBarony struct {
	feud.BaseBarony
}

var BEngelberg恩格尔贝格 feud.Barony = &恩格尔贝格EngelbergBarony{}

func init() {
    f := BEngelberg恩格尔贝格.(*恩格尔贝格EngelbergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "engelberg",
		TitleName: "恩格尔贝格",
		TitleCode: "b_engelberg",
	}
}
