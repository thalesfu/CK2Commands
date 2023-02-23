package holland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 豪达GoudaBarony struct {
	feud.BaseBarony
}

var BGouda豪达 feud.Barony = &豪达GoudaBarony{}

func init() {
	f := BGouda豪达.(*豪达GoudaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gouda",
		TitleName: "豪达",
		TitleCode: "b_gouda",
	}
}
