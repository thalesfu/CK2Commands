package ajadabiya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凯尔萨KarsahBarony struct {
	feud.BaseBarony
}

var BKarsah凯尔萨 feud.Barony = &凯尔萨KarsahBarony{}

func init() {
    f := BKarsah凯尔萨.(*凯尔萨KarsahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karsah",
		TitleName: "凯尔萨",
		TitleCode: "b_karsah",
	}
}
