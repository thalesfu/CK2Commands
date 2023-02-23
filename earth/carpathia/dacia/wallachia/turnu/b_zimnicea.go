package turnu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 济姆尼恰ZimniceaBarony struct {
	feud.BaseBarony
}

var BZimnicea济姆尼恰 feud.Barony = &济姆尼恰ZimniceaBarony{}

func init() {
	f := BZimnicea济姆尼恰.(*济姆尼恰ZimniceaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zimnicea",
		TitleName: "济姆尼恰",
		TitleCode: "b_zimnicea",
	}
}
