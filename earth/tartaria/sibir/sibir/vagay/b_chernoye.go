package vagay

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 切尔诺耶ChernoyeBarony struct {
	feud.BaseBarony
}

var BChernoye切尔诺耶 feud.Barony = &切尔诺耶ChernoyeBarony{}

func init() {
    f := BChernoye切尔诺耶.(*切尔诺耶ChernoyeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chernoye",
		TitleName: "切尔诺耶",
		TitleCode: "b_chernoye",
	}
}
