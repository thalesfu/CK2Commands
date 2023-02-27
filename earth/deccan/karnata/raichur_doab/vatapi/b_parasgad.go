package vatapi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕拉斯迦得ParasgadBarony struct {
	feud.BaseBarony
}

var BParasgad帕拉斯迦得 feud.Barony = &帕拉斯迦得ParasgadBarony{}

func init() {
    f := BParasgad帕拉斯迦得.(*帕拉斯迦得ParasgadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "parasgad",
		TitleName: "帕拉斯迦得",
		TitleCode: "b_parasgad",
	}
}
