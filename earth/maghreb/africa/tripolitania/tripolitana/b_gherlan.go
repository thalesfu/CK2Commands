package tripolitana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格赫尔兰GherlanBarony struct {
	feud.BaseBarony
}

var BGherlan格赫尔兰 feud.Barony = &格赫尔兰GherlanBarony{}

func init() {
	f := BGherlan格赫尔兰.(*格赫尔兰GherlanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gherlan",
		TitleName: "格赫尔兰",
		TitleCode: "b_gherlan",
	}
}
