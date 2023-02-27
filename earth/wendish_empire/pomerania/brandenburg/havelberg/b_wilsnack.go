package havelberg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维尔斯纳克WilsnackBarony struct {
	feud.BaseBarony
}

var BWilsnack维尔斯纳克 feud.Barony = &维尔斯纳克WilsnackBarony{}

func init() {
    f := BWilsnack维尔斯纳克.(*维尔斯纳克WilsnackBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wilsnack",
		TitleName: "维尔斯纳克",
		TitleCode: "b_wilsnack",
	}
}
