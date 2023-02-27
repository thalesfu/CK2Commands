package lugo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳维卢比奥NavilubioBarony struct {
	feud.BaseBarony
}

var BNavilubio纳维卢比奥 feud.Barony = &纳维卢比奥NavilubioBarony{}

func init() {
    f := BNavilubio纳维卢比奥.(*纳维卢比奥NavilubioBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "navilubio",
		TitleName: "纳维卢比奥",
		TitleCode: "b_navilubio",
	}
}
