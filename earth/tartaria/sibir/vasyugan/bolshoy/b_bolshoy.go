package bolshoy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 大尤甘BolshoyBarony struct {
	feud.BaseBarony
}

var BBolshoy大尤甘 feud.Barony = &大尤甘BolshoyBarony{}

func init() {
    f := BBolshoy大尤甘.(*大尤甘BolshoyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bolshoy",
		TitleName: "大尤甘",
		TitleCode: "b_bolshoy",
	}
}
