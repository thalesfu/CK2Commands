package aral

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿谢_阿特里克AszczeatrikBarony struct {
	feud.BaseBarony
}

var BAszczeatrik阿谢_阿特里克 feud.Barony = &阿谢_阿特里克AszczeatrikBarony{}

func init() {
    f := BAszczeatrik阿谢_阿特里克.(*阿谢_阿特里克AszczeatrikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aszczeatrik",
		TitleName: "阿谢_阿特里克",
		TitleCode: "b_aszczeatrik",
	}
}
