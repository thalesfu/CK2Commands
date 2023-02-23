package thessalonike

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞尔维亚ServiaBarony struct {
	feud.BaseBarony
}

var BServia塞尔维亚 feud.Barony = &塞尔维亚ServiaBarony{}

func init() {
	f := BServia塞尔维亚.(*塞尔维亚ServiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "servia",
		TitleName: "塞尔维亚",
		TitleCode: "b_servia",
	}
}
