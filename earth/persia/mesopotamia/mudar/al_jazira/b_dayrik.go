package al_jazira

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 代里克DayrikBarony struct {
	feud.BaseBarony
}

var BDayrik代里克 feud.Barony = &代里克DayrikBarony{}

func init() {
    f := BDayrik代里克.(*代里克DayrikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dayrik",
		TitleName: "代里克",
		TitleCode: "b_dayrik",
	}
}
