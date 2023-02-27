package muscat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马斯喀特MuscatBarony struct {
	feud.BaseBarony
}

var BMuscat马斯喀特 feud.Barony = &马斯喀特MuscatBarony{}

func init() {
    f := BMuscat马斯喀特.(*马斯喀特MuscatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "muscat",
		TitleName: "马斯喀特",
		TitleCode: "b_muscat",
	}
}
