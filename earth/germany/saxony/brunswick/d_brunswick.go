package brunswick

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BrunswickDuke interface {
    feud.Duke
}

type 布伦瑞克BrunswickDuke struct {
	feud.BaseDuke
}

var DBrunswick布伦瑞克 BrunswickDuke = &布伦瑞克BrunswickDuke{}

func init() {
	f := DBrunswick布伦瑞克.(*布伦瑞克BrunswickDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "brunswick",
		TitleName: "布伦瑞克",
		TitleCode: "d_brunswick",
		Counties:  map[string]feud.County{},
	}

}
