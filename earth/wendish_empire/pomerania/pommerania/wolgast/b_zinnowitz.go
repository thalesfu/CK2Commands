package wolgast

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 钦诺维茨ZinnowitzBarony struct {
	feud.BaseBarony
}

var BZinnowitz钦诺维茨 feud.Barony = &钦诺维茨ZinnowitzBarony{}

func init() {
    f := BZinnowitz钦诺维茨.(*钦诺维茨ZinnowitzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zinnowitz",
		TitleName: "钦诺维茨",
		TitleCode: "b_zinnowitz",
	}
}
