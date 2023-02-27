package syrj

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 热沙尔特ZheshartBarony struct {
	feud.BaseBarony
}

var BZheshart热沙尔特 feud.Barony = &热沙尔特ZheshartBarony{}

func init() {
    f := BZheshart热沙尔特.(*热沙尔特ZheshartBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zheshart",
		TitleName: "热沙尔特",
		TitleCode: "b_zheshart",
	}
}
