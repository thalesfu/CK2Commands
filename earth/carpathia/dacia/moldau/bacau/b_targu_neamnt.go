package bacau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特尔古尼亚姆茨Targu_neamntBarony struct {
	feud.BaseBarony
}

var BTargu_neamnt特尔古尼亚姆茨 feud.Barony = &特尔古尼亚姆茨Targu_neamntBarony{}

func init() {
    f := BTargu_neamnt特尔古尼亚姆茨.(*特尔古尼亚姆茨Targu_neamntBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "targu_neamnt",
		TitleName: "特尔古尼亚姆茨",
		TitleCode: "b_targu_neamnt",
	}
}
