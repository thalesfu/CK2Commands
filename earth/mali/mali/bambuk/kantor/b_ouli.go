package kantor

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌利OuliBarony struct {
	feud.BaseBarony
}

var BOuli乌利 feud.Barony = &乌利OuliBarony{}

func init() {
	f := BOuli乌利.(*乌利OuliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ouli",
		TitleName: "乌利",
		TitleCode: "b_ouli",
	}
}
