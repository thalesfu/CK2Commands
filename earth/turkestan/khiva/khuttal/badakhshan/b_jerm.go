package badakhshan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杰姆JermBarony struct {
	feud.BaseBarony
}

var BJerm杰姆 feud.Barony = &杰姆JermBarony{}

func init() {
	f := BJerm杰姆.(*杰姆JermBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jerm",
		TitleName: "杰姆",
		TitleCode: "b_jerm",
	}
}
