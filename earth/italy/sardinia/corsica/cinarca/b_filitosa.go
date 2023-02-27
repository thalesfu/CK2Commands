package cinarca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 菲利托萨FilitosaBarony struct {
	feud.BaseBarony
}

var BFilitosa菲利托萨 feud.Barony = &菲利托萨FilitosaBarony{}

func init() {
    f := BFilitosa菲利托萨.(*菲利托萨FilitosaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "filitosa",
		TitleName: "菲利托萨",
		TitleCode: "b_filitosa",
	}
}
