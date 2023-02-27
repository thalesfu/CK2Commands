package havelberg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈弗尔贝格HavelbergBarony struct {
	feud.BaseBarony
}

var BHavelberg哈弗尔贝格 feud.Barony = &哈弗尔贝格HavelbergBarony{}

func init() {
    f := BHavelberg哈弗尔贝格.(*哈弗尔贝格HavelbergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "havelberg",
		TitleName: "哈弗尔贝格",
		TitleCode: "b_havelberg",
	}
}
