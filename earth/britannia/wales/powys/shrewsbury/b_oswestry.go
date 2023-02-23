package shrewsbury

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥斯沃斯特里OswestryBarony struct {
	feud.BaseBarony
}

var BOswestry奥斯沃斯特里 feud.Barony = &奥斯沃斯特里OswestryBarony{}

func init() {
	f := BOswestry奥斯沃斯特里.(*奥斯沃斯特里OswestryBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oswestry",
		TitleName: "奥斯沃斯特里",
		TitleCode: "b_oswestry",
	}
}
