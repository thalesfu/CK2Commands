package orekhovo_zouievo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 叶戈里耶夫斯克YegoryevskBarony struct {
	feud.BaseBarony
}

var BYegoryevsk叶戈里耶夫斯克 feud.Barony = &叶戈里耶夫斯克YegoryevskBarony{}

func init() {
    f := BYegoryevsk叶戈里耶夫斯克.(*叶戈里耶夫斯克YegoryevskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yegoryevsk",
		TitleName: "叶戈里耶夫斯克",
		TitleCode: "b_yegoryevsk",
	}
}
