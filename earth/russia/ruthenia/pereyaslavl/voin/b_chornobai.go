package voin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乔尔诺拜ChornobaiBarony struct {
	feud.BaseBarony
}

var BChornobai乔尔诺拜 feud.Barony = &乔尔诺拜ChornobaiBarony{}

func init() {
    f := BChornobai乔尔诺拜.(*乔尔诺拜ChornobaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chornobai",
		TitleName: "乔尔诺拜",
		TitleCode: "b_chornobai",
	}
}
