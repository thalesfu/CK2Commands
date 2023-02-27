package taghaza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿乌格鲁特AougroutBarony struct {
	feud.BaseBarony
}

var BAougrout阿乌格鲁特 feud.Barony = &阿乌格鲁特AougroutBarony{}

func init() {
    f := BAougrout阿乌格鲁特.(*阿乌格鲁特AougroutBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aougrout",
		TitleName: "阿乌格鲁特",
		TitleCode: "b_aougrout",
	}
}
