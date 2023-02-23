package alexandretta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗特拉PortellaBarony struct {
	feud.BaseBarony
}

var BPortella罗特拉 feud.Barony = &罗特拉PortellaBarony{}

func init() {
	f := BPortella罗特拉.(*罗特拉PortellaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "portella",
		TitleName: "罗特拉",
		TitleCode: "b_portella",
	}
}
