package balkonda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 夏罗布特KhairaputBarony struct {
	feud.BaseBarony
}

var BKhairaput夏罗布特 feud.Barony = &夏罗布特KhairaputBarony{}

func init() {
    f := BKhairaput夏罗布特.(*夏罗布特KhairaputBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khairaput",
		TitleName: "夏罗布特",
		TitleCode: "b_khairaput",
	}
}
