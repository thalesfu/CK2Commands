package alexandretta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博内勒港PortbonnelBarony struct {
	feud.BaseBarony
}

var BPortbonnel博内勒港 feud.Barony = &博内勒港PortbonnelBarony{}

func init() {
	f := BPortbonnel博内勒港.(*博内勒港PortbonnelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "portbonnel",
		TitleName: "博内勒港",
		TitleCode: "b_portbonnel",
	}
}
