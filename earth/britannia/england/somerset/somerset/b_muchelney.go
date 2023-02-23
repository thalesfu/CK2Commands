package somerset

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马乔尼MuchelneyBarony struct {
	feud.BaseBarony
}

var BMuchelney马乔尼 feud.Barony = &马乔尼MuchelneyBarony{}

func init() {
	f := BMuchelney马乔尼.(*马乔尼MuchelneyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "muchelney",
		TitleName: "马乔尼",
		TitleCode: "b_muchelney",
	}
}
