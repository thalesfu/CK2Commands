package gemer

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 约尔什沃JolsvaBarony struct {
	feud.BaseBarony
}

var BJolsva约尔什沃 feud.Barony = &约尔什沃JolsvaBarony{}

func init() {
	f := BJolsva约尔什沃.(*约尔什沃JolsvaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jolsva",
		TitleName: "约尔什沃",
		TitleCode: "b_jolsva",
	}
}
