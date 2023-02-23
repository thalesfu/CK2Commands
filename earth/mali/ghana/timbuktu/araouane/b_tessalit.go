package araouane

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泰萨利特TessalitBarony struct {
	feud.BaseBarony
}

var BTessalit泰萨利特 feud.Barony = &泰萨利特TessalitBarony{}

func init() {
	f := BTessalit泰萨利特.(*泰萨利特TessalitBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tessalit",
		TitleName: "泰萨利特",
		TitleCode: "b_tessalit",
	}
}
