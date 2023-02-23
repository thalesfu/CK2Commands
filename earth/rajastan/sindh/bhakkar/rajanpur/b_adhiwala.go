package rajanpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 代希瓦勒AdhiwalaBarony struct {
	feud.BaseBarony
}

var BAdhiwala代希瓦勒 feud.Barony = &代希瓦勒AdhiwalaBarony{}

func init() {
	f := BAdhiwala代希瓦勒.(*代希瓦勒AdhiwalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "adhiwala",
		TitleName: "代希瓦勒",
		TitleCode: "b_adhiwala",
	}
}
