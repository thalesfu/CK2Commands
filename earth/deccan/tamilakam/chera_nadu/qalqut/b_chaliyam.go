package qalqut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伽利耶ChaliyamBarony struct {
	feud.BaseBarony
}

var BChaliyam伽利耶 feud.Barony = &伽利耶ChaliyamBarony{}

func init() {
    f := BChaliyam伽利耶.(*伽利耶ChaliyamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chaliyam",
		TitleName: "伽利耶",
		TitleCode: "b_chaliyam",
	}
}
