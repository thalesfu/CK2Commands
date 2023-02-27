package sedyu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尤瓦Ayuva_sedyuBarony struct {
	feud.BaseBarony
}

var BAyuva_sedyu阿尤瓦 feud.Barony = &阿尤瓦Ayuva_sedyuBarony{}

func init() {
    f := BAyuva_sedyu阿尤瓦.(*阿尤瓦Ayuva_sedyuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ayuva_sedyu",
		TitleName: "阿尤瓦",
		TitleCode: "b_ayuva_sedyu",
	}
}
