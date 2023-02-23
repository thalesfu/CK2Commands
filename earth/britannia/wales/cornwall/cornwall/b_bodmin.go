package cornwall

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博德明BodminBarony struct {
	feud.BaseBarony
}

var BBodmin博德明 feud.Barony = &博德明BodminBarony{}

func init() {
	f := BBodmin博德明.(*博德明BodminBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bodmin",
		TitleName: "博德明",
		TitleCode: "b_bodmin",
	}
}
