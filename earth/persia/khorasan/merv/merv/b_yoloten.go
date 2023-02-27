package merv

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 约廖坚YolotenBarony struct {
	feud.BaseBarony
}

var BYoloten约廖坚 feud.Barony = &约廖坚YolotenBarony{}

func init() {
    f := BYoloten约廖坚.(*约廖坚YolotenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yoloten",
		TitleName: "约廖坚",
		TitleCode: "b_yoloten",
	}
}
