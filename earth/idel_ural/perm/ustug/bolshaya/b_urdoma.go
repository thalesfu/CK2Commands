package bolshaya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌尔多马UrdomaBarony struct {
	feud.BaseBarony
}

var BUrdoma乌尔多马 feud.Barony = &乌尔多马UrdomaBarony{}

func init() {
    f := BUrdoma乌尔多马.(*乌尔多马UrdomaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "urdoma",
		TitleName: "乌尔多马",
		TitleCode: "b_urdoma",
	}
}
