package brabant

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布鲁塞尔BrusselBarony struct {
	feud.BaseBarony
}

var BBrussel布鲁塞尔 feud.Barony = &布鲁塞尔BrusselBarony{}

func init() {
	f := BBrussel布鲁塞尔.(*布鲁塞尔BrusselBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "brussel",
		TitleName: "布鲁塞尔",
		TitleCode: "b_brussel",
	}
}
