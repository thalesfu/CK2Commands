package aror

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿卢梨ArorBarony struct {
	feud.BaseBarony
}

var BAror阿卢梨 feud.Barony = &阿卢梨ArorBarony{}

func init() {
    f := BAror阿卢梨.(*阿卢梨ArorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aror",
		TitleName: "阿卢梨",
		TitleCode: "b_aror",
	}
}
