package muztau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 济良ZyryanBarony struct {
	feud.BaseBarony
}

var BZyryan济良 feud.Barony = &济良ZyryanBarony{}

func init() {
    f := BZyryan济良.(*济良ZyryanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zyryan",
		TitleName: "济良",
		TitleCode: "b_zyryan",
	}
}
