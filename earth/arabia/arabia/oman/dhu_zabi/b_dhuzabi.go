package dhu_zabi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿布扎比DhuzabiBarony struct {
	feud.BaseBarony
}

var BDhuzabi阿布扎比 feud.Barony = &阿布扎比DhuzabiBarony{}

func init() {
    f := BDhuzabi阿布扎比.(*阿布扎比DhuzabiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dhuzabi",
		TitleName: "阿布扎比",
		TitleCode: "b_dhuzabi",
	}
}
