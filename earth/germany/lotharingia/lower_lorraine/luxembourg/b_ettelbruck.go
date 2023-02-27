package luxembourg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃特尔布吕克EttelbruckBarony struct {
	feud.BaseBarony
}

var BEttelbruck埃特尔布吕克 feud.Barony = &埃特尔布吕克EttelbruckBarony{}

func init() {
    f := BEttelbruck埃特尔布吕克.(*埃特尔布吕克EttelbruckBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ettelbruck",
		TitleName: "埃特尔布吕克",
		TitleCode: "b_ettelbruck",
	}
}
