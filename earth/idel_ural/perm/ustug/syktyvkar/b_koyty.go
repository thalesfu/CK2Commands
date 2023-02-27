package syktyvkar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科伊特KoytyBarony struct {
	feud.BaseBarony
}

var BKoyty科伊特 feud.Barony = &科伊特KoytyBarony{}

func init() {
    f := BKoyty科伊特.(*科伊特KoytyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "koyty",
		TitleName: "科伊特",
		TitleCode: "b_koyty",
	}
}
