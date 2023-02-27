package dimapur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尸拘吒AskotBarony struct {
	feud.BaseBarony
}

var BAskot阿尸拘吒 feud.Barony = &阿尸拘吒AskotBarony{}

func init() {
    f := BAskot阿尸拘吒.(*阿尸拘吒AskotBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "askot",
		TitleName: "阿尸拘吒",
		TitleCode: "b_askot",
	}
}
