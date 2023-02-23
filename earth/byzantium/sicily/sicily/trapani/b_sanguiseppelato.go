package trapani

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣朱塞佩亚托SanguiseppelatoBarony struct {
	feud.BaseBarony
}

var BSanguiseppelato圣朱塞佩亚托 feud.Barony = &圣朱塞佩亚托SanguiseppelatoBarony{}

func init() {
	f := BSanguiseppelato圣朱塞佩亚托.(*圣朱塞佩亚托SanguiseppelatoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sanguiseppelato",
		TitleName: "圣朱塞佩亚托",
		TitleCode: "b_sanguiseppelato",
	}
}
