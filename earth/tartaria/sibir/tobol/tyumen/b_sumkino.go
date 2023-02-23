package tyumen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏姆基诺SumkinoBarony struct {
	feud.BaseBarony
}

var BSumkino苏姆基诺 feud.Barony = &苏姆基诺SumkinoBarony{}

func init() {
	f := BSumkino苏姆基诺.(*苏姆基诺SumkinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sumkino",
		TitleName: "苏姆基诺",
		TitleCode: "b_sumkino",
	}
}
