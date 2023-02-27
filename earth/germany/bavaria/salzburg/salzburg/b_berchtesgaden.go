package salzburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝希特斯加登BerchtesgadenBarony struct {
	feud.BaseBarony
}

var BBerchtesgaden贝希特斯加登 feud.Barony = &贝希特斯加登BerchtesgadenBarony{}

func init() {
    f := BBerchtesgaden贝希特斯加登.(*贝希特斯加登BerchtesgadenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "berchtesgaden",
		TitleName: "贝希特斯加登",
		TitleCode: "b_berchtesgaden",
	}
}
