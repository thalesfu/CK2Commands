package lakhnau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阇耶陀城ZaidpurBarony struct {
	feud.BaseBarony
}

var BZaidpur阇耶陀城 feud.Barony = &阇耶陀城ZaidpurBarony{}

func init() {
    f := BZaidpur阇耶陀城.(*阇耶陀城ZaidpurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zaidpur",
		TitleName: "阇耶陀城",
		TitleCode: "b_zaidpur",
	}
}
