package koln

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布劳韦勒BrauweilerBarony struct {
	feud.BaseBarony
}

var BBrauweiler布劳韦勒 feud.Barony = &布劳韦勒BrauweilerBarony{}

func init() {
	f := BBrauweiler布劳韦勒.(*布劳韦勒BrauweilerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "brauweiler",
		TitleName: "布劳韦勒",
		TitleCode: "b_brauweiler",
	}
}
