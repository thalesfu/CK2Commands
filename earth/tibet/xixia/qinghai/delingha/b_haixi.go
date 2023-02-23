package delingha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 海西HaixiBarony struct {
	feud.BaseBarony
}

var BHaixi海西 feud.Barony = &海西HaixiBarony{}

func init() {
	f := BHaixi海西.(*海西HaixiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "haixi",
		TitleName: "海西",
		TitleCode: "b_haixi",
	}
}
