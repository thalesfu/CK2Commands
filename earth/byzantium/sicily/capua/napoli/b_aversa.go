package napoli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿韦尔萨AversaBarony struct {
	feud.BaseBarony
}

var BAversa阿韦尔萨 feud.Barony = &阿韦尔萨AversaBarony{}

func init() {
	f := BAversa阿韦尔萨.(*阿韦尔萨AversaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aversa",
		TitleName: "阿韦尔萨",
		TitleCode: "b_aversa",
	}
}
