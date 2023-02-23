package tavasts

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦纳亚VanajaBarony struct {
	feud.BaseBarony
}

var BVanaja瓦纳亚 feud.Barony = &瓦纳亚VanajaBarony{}

func init() {
	f := BVanaja瓦纳亚.(*瓦纳亚VanajaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vanaja",
		TitleName: "瓦纳亚",
		TitleCode: "b_vanaja",
	}
}
