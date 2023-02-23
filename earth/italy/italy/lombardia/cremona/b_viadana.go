package cremona

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维亚达纳ViadanaBarony struct {
	feud.BaseBarony
}

var BViadana维亚达纳 feud.Barony = &维亚达纳ViadanaBarony{}

func init() {
	f := BViadana维亚达纳.(*维亚达纳ViadanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "viadana",
		TitleName: "维亚达纳",
		TitleCode: "b_viadana",
	}
}
