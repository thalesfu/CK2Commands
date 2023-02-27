package korchev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 切尔格ChercoBarony struct {
	feud.BaseBarony
}

var BCherco切尔格 feud.Barony = &切尔格ChercoBarony{}

func init() {
    f := BCherco切尔格.(*切尔格ChercoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cherco",
		TitleName: "切尔格",
		TitleCode: "b_cherco",
	}
}
