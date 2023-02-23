package pratishthana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 平格尔BhingarBarony struct {
	feud.BaseBarony
}

var BBhingar平格尔 feud.Barony = &平格尔BhingarBarony{}

func init() {
	f := BBhingar平格尔.(*平格尔BhingarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bhingar",
		TitleName: "平格尔",
		TitleCode: "b_bhingar",
	}
}
