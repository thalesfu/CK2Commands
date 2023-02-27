package uvek

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马克思MarksBarony struct {
	feud.BaseBarony
}

var BMarks马克思 feud.Barony = &马克思MarksBarony{}

func init() {
    f := BMarks马克思.(*马克思MarksBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "marks",
		TitleName: "马克思",
		TitleCode: "b_marks",
	}
}
