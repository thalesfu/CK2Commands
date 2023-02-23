package herjedalen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦姆达伦VemdalenBarony struct {
	feud.BaseBarony
}

var BVemdalen韦姆达伦 feud.Barony = &韦姆达伦VemdalenBarony{}

func init() {
	f := BVemdalen韦姆达伦.(*韦姆达伦VemdalenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vemdalen",
		TitleName: "韦姆达伦",
		TitleCode: "b_vemdalen",
	}
}
