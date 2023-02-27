package votyaki

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比谢尔季BisertBarony struct {
	feud.BaseBarony
}

var BBisert比谢尔季 feud.Barony = &比谢尔季BisertBarony{}

func init() {
    f := BBisert比谢尔季.(*比谢尔季BisertBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bisert",
		TitleName: "比谢尔季",
		TitleCode: "b_bisert",
	}
}
