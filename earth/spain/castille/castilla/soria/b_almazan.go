package soria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔马桑AlmazanBarony struct {
	feud.BaseBarony
}

var BAlmazan阿尔马桑 feud.Barony = &阿尔马桑AlmazanBarony{}

func init() {
    f := BAlmazan阿尔马桑.(*阿尔马桑AlmazanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "almazan",
		TitleName: "阿尔马桑",
		TitleCode: "b_almazan",
	}
}
