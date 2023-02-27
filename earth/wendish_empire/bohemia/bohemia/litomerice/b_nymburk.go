package litomerice

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 宁布尔克NymburkBarony struct {
	feud.BaseBarony
}

var BNymburk宁布尔克 feud.Barony = &宁布尔克NymburkBarony{}

func init() {
    f := BNymburk宁布尔克.(*宁布尔克NymburkBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nymburk",
		TitleName: "宁布尔克",
		TitleCode: "b_nymburk",
	}
}
