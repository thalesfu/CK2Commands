package aoudaghost

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥达戈斯特AoudaghostBarony struct {
	feud.BaseBarony
}

var BAoudaghost奥达戈斯特 feud.Barony = &奥达戈斯特AoudaghostBarony{}

func init() {
    f := BAoudaghost奥达戈斯特.(*奥达戈斯特AoudaghostBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aoudaghost",
		TitleName: "奥达戈斯特",
		TitleCode: "b_aoudaghost",
	}
}
