package shirvan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希济尔津达KhizirzindaBarony struct {
	feud.BaseBarony
}

var BKhizirzinda希济尔津达 feud.Barony = &希济尔津达KhizirzindaBarony{}

func init() {
    f := BKhizirzinda希济尔津达.(*希济尔津达KhizirzindaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khizirzinda",
		TitleName: "希济尔津达",
		TitleCode: "b_khizirzinda",
	}
}
