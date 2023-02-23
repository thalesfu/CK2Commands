package devon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利德福德LydfordBarony struct {
	feud.BaseBarony
}

var BLydford利德福德 feud.Barony = &利德福德LydfordBarony{}

func init() {
	f := BLydford利德福德.(*利德福德LydfordBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lydford",
		TitleName: "利德福德",
		TitleCode: "b_lydford",
	}
}
