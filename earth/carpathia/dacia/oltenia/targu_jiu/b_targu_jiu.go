package targu_jiu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特尔古日乌Targu_jiuBarony struct {
	feud.BaseBarony
}

var BTargu_jiu特尔古日乌 feud.Barony = &特尔古日乌Targu_jiuBarony{}

func init() {
    f := BTargu_jiu特尔古日乌.(*特尔古日乌Targu_jiuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "targu_jiu",
		TitleName: "特尔古日乌",
		TitleCode: "b_targu_jiu",
	}
}
