package gottingen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 柯维CorveyBarony struct {
	feud.BaseBarony
}

var BCorvey柯维 feud.Barony = &柯维CorveyBarony{}

func init() {
    f := BCorvey柯维.(*柯维CorveyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "corvey",
		TitleName: "柯维",
		TitleCode: "b_corvey",
	}
}
