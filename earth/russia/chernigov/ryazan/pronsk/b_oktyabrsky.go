package pronsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 十月镇OktyabrskyBarony struct {
	feud.BaseBarony
}

var BOktyabrsky十月镇 feud.Barony = &十月镇OktyabrskyBarony{}

func init() {
	f := BOktyabrsky十月镇.(*十月镇OktyabrskyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oktyabrsky",
		TitleName: "十月镇",
		TitleCode: "b_oktyabrsky",
	}
}
