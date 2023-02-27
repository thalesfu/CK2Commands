package kondana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 毗摩商羯罗BhimashankaraBarony struct {
	feud.BaseBarony
}

var BBhimashankara毗摩商羯罗 feud.Barony = &毗摩商羯罗BhimashankaraBarony{}

func init() {
    f := BBhimashankara毗摩商羯罗.(*毗摩商羯罗BhimashankaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bhimashankara",
		TitleName: "毗摩商羯罗",
		TitleCode: "b_bhimashankara",
	}
}
