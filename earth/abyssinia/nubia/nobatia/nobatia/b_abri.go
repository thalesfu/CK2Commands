package nobatia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿卜里AbriBarony struct {
	feud.BaseBarony
}

var BAbri阿卜里 feud.Barony = &阿卜里AbriBarony{}

func init() {
	f := BAbri阿卜里.(*阿卜里AbriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abri",
		TitleName: "阿卜里",
		TitleCode: "b_abri",
	}
}
