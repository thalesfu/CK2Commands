package adrianopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迪蒂莫迪索DidymoteichonBarony struct {
	feud.BaseBarony
}

var BDidymoteichon迪蒂莫迪索 feud.Barony = &迪蒂莫迪索DidymoteichonBarony{}

func init() {
    f := BDidymoteichon迪蒂莫迪索.(*迪蒂莫迪索DidymoteichonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "didymoteichon",
		TitleName: "迪蒂莫迪索",
		TitleCode: "b_didymoteichon",
	}
}
