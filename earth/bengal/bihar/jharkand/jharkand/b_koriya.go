package jharkand

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 憍利耶KoriyaBarony struct {
	feud.BaseBarony
}

var BKoriya憍利耶 feud.Barony = &憍利耶KoriyaBarony{}

func init() {
    f := BKoriya憍利耶.(*憍利耶KoriyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "koriya",
		TitleName: "憍利耶",
		TitleCode: "b_koriya",
	}
}
