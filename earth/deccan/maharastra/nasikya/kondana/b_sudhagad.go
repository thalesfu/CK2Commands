package kondana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏陀迦荼SudhagadBarony struct {
	feud.BaseBarony
}

var BSudhagad苏陀迦荼 feud.Barony = &苏陀迦荼SudhagadBarony{}

func init() {
	f := BSudhagad苏陀迦荼.(*苏陀迦荼SudhagadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sudhagad",
		TitleName: "苏陀迦荼",
		TitleCode: "b_sudhagad",
	}
}
