package aylik

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌苏UsuBarony struct {
	feud.BaseBarony
}

var BUsu乌苏 feud.Barony = &乌苏UsuBarony{}

func init() {
    f := BUsu乌苏.(*乌苏UsuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "usu",
		TitleName: "乌苏",
		TitleCode: "b_usu",
	}
}
