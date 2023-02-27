package hisar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿迦卢诃AgrohaBarony struct {
	feud.BaseBarony
}

var BAgroha阿迦卢诃 feud.Barony = &阿迦卢诃AgrohaBarony{}

func init() {
    f := BAgroha阿迦卢诃.(*阿迦卢诃AgrohaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "agroha",
		TitleName: "阿迦卢诃",
		TitleCode: "b_agroha",
	}
}
