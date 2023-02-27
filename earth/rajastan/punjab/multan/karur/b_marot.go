package karur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 默罗德MarotBarony struct {
	feud.BaseBarony
}

var BMarot默罗德 feud.Barony = &默罗德MarotBarony{}

func init() {
    f := BMarot默罗德.(*默罗德MarotBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "marot",
		TitleName: "默罗德",
		TitleCode: "b_marot",
	}
}
