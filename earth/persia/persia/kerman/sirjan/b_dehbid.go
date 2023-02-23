package sirjan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 代比德DehbidBarony struct {
	feud.BaseBarony
}

var BDehbid代比德 feud.Barony = &代比德DehbidBarony{}

func init() {
	f := BDehbid代比德.(*代比德DehbidBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dehbid",
		TitleName: "代比德",
		TitleCode: "b_dehbid",
	}
}
