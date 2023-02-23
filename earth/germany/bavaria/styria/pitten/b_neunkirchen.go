package pitten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诺因基兴NeunkirchenBarony struct {
	feud.BaseBarony
}

var BNeunkirchen诺因基兴 feud.Barony = &诺因基兴NeunkirchenBarony{}

func init() {
	f := BNeunkirchen诺因基兴.(*诺因基兴NeunkirchenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "neunkirchen",
		TitleName: "诺因基兴",
		TitleCode: "b_neunkirchen",
	}
}
