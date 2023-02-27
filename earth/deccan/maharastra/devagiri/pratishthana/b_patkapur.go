package pratishthana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕提加普尔PatkapurBarony struct {
	feud.BaseBarony
}

var BPatkapur帕提加普尔 feud.Barony = &帕提加普尔PatkapurBarony{}

func init() {
    f := BPatkapur帕提加普尔.(*帕提加普尔PatkapurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "patkapur",
		TitleName: "帕提加普尔",
		TitleCode: "b_patkapur",
	}
}
