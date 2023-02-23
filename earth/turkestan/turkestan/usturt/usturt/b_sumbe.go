package usturt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 孙别SumbeBarony struct {
	feud.BaseBarony
}

var BSumbe孙别 feud.Barony = &孙别SumbeBarony{}

func init() {
	f := BSumbe孙别.(*孙别SumbeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sumbe",
		TitleName: "孙别",
		TitleCode: "b_sumbe",
	}
}
