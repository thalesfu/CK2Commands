package bratslav

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托马什皮尔TomashpilBarony struct {
	feud.BaseBarony
}

var BTomashpil托马什皮尔 feud.Barony = &托马什皮尔TomashpilBarony{}

func init() {
    f := BTomashpil托马什皮尔.(*托马什皮尔TomashpilBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tomashpil",
		TitleName: "托马什皮尔",
		TitleCode: "b_tomashpil",
	}
}
