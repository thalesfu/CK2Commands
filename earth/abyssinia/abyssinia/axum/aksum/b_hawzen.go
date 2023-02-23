package aksum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 豪赞HawzenBarony struct {
	feud.BaseBarony
}

var BHawzen豪赞 feud.Barony = &豪赞HawzenBarony{}

func init() {
	f := BHawzen豪赞.(*豪赞HawzenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hawzen",
		TitleName: "豪赞",
		TitleCode: "b_hawzen",
	}
}
