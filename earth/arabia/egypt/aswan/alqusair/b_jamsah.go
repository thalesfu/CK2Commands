package alqusair

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杰姆萨JamsahBarony struct {
	feud.BaseBarony
}

var BJamsah杰姆萨 feud.Barony = &杰姆萨JamsahBarony{}

func init() {
    f := BJamsah杰姆萨.(*杰姆萨JamsahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jamsah",
		TitleName: "杰姆萨",
		TitleCode: "b_jamsah",
	}
}
