package lausitz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卢考LuckauBarony struct {
	feud.BaseBarony
}

var BLuckau卢考 feud.Barony = &卢考LuckauBarony{}

func init() {
    f := BLuckau卢考.(*卢考LuckauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "luckau",
		TitleName: "卢考",
		TitleCode: "b_luckau",
	}
}
