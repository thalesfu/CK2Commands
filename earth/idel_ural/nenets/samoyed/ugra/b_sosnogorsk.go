package ugra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索斯诺戈尔斯克SosnogorskBarony struct {
	feud.BaseBarony
}

var BSosnogorsk索斯诺戈尔斯克 feud.Barony = &索斯诺戈尔斯克SosnogorskBarony{}

func init() {
    f := BSosnogorsk索斯诺戈尔斯克.(*索斯诺戈尔斯克SosnogorskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sosnogorsk",
		TitleName: "索斯诺戈尔斯克",
		TitleCode: "b_sosnogorsk",
	}
}
