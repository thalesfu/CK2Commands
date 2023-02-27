package chunar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 周那罗ChunarBarony struct {
	feud.BaseBarony
}

var BChunar周那罗 feud.Barony = &周那罗ChunarBarony{}

func init() {
    f := BChunar周那罗.(*周那罗ChunarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chunar",
		TitleName: "周那罗",
		TitleCode: "b_chunar",
	}
}
