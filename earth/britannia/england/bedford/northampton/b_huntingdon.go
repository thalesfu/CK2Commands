package northampton

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亨廷登HuntingdonBarony struct {
	feud.BaseBarony
}

var BHuntingdon亨廷登 feud.Barony = &亨廷登HuntingdonBarony{}

func init() {
    f := BHuntingdon亨廷登.(*亨廷登HuntingdonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "huntingdon",
		TitleName: "亨廷登",
		TitleCode: "b_huntingdon",
	}
}
