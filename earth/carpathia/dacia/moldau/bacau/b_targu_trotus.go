package bacau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特尔古特罗图什Targu_trotusBarony struct {
	feud.BaseBarony
}

var BTargu_trotus特尔古特罗图什 feud.Barony = &特尔古特罗图什Targu_trotusBarony{}

func init() {
    f := BTargu_trotus特尔古特罗图什.(*特尔古特罗图什Targu_trotusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "targu_trotus",
		TitleName: "特尔古特罗图什",
		TitleCode: "b_targu_trotus",
	}
}
