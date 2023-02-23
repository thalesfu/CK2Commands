package ujjayini

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩醯陀补罗MehidpurBarony struct {
	feud.BaseBarony
}

var BMehidpur摩醯陀补罗 feud.Barony = &摩醯陀补罗MehidpurBarony{}

func init() {
	f := BMehidpur摩醯陀补罗.(*摩醯陀补罗MehidpurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mehidpur",
		TitleName: "摩醯陀补罗",
		TitleCode: "b_mehidpur",
	}
}
