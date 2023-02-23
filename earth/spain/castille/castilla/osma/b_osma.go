package osma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥斯马OsmaBarony struct {
	feud.BaseBarony
}

var BOsma奥斯马 feud.Barony = &奥斯马OsmaBarony{}

func init() {
	f := BOsma奥斯马.(*奥斯马OsmaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "osma",
		TitleName: "奥斯马",
		TitleCode: "b_osma",
	}
}
