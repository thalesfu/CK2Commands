package tell_atlas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔古格哈TagoughaltBarony struct {
	feud.BaseBarony
}

var BTagoughalt塔古格哈 feud.Barony = &塔古格哈TagoughaltBarony{}

func init() {
    f := BTagoughalt塔古格哈.(*塔古格哈TagoughaltBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tagoughalt",
		TitleName: "塔古格哈",
		TitleCode: "b_tagoughalt",
	}
}
