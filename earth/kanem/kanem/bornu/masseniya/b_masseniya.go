package masseniya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马塞尼亚MasseniyaBarony struct {
	feud.BaseBarony
}

var BMasseniya马塞尼亚 feud.Barony = &马塞尼亚MasseniyaBarony{}

func init() {
	f := BMasseniya马塞尼亚.(*马塞尼亚MasseniyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "masseniya",
		TitleName: "马塞尼亚",
		TitleCode: "b_masseniya",
	}
}
