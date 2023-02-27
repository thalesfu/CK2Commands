package udmurts

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃特金斯克VotkinskBarony struct {
	feud.BaseBarony
}

var BVotkinsk沃特金斯克 feud.Barony = &沃特金斯克VotkinskBarony{}

func init() {
    f := BVotkinsk沃特金斯克.(*沃特金斯克VotkinskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "votkinsk",
		TitleName: "沃特金斯克",
		TitleCode: "b_votkinsk",
	}
}
