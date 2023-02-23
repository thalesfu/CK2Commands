package afghanistan

import (
	"github.com/thalesfu/CK2Commands/earth/persia/afghanistan/kabul"
	"github.com/thalesfu/CK2Commands/earth/persia/afghanistan/zabulistan"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AfghanistanKingdom interface {
	feud.Kingdom
	DKabul迦布罗() kabul.KabulDuke
	DZabulistan社护罗萨他那() zabulistan.ZabulistanDuke
}

type 迦布罗AfghanistanKingdom struct {
	feud.BaseKingdom
	迦布罗Kabul         kabul.KabulDuke
	社护罗萨他那Zabulistan zabulistan.ZabulistanDuke
}

func (k *迦布罗AfghanistanKingdom) DKabul迦布罗() kabul.KabulDuke {
	return k.迦布罗Kabul
}

func (k *迦布罗AfghanistanKingdom) DZabulistan社护罗萨他那() zabulistan.ZabulistanDuke {
	return k.社护罗萨他那Zabulistan
}

var KAfghanistan迦布罗 AfghanistanKingdom = &迦布罗AfghanistanKingdom{}

func init() {
	f := KAfghanistan迦布罗.(*迦布罗AfghanistanKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "afghanistan",
		TitleName: "迦布罗",
		TitleCode: "k_afghanistan",
		Dukes:     map[string]feud.Duke{},
	}

	f.迦布罗Kabul = kabul.DKabul迦布罗
	f.迦布罗Kabul.SetParent(f)

	f.社护罗萨他那Zabulistan = zabulistan.DZabulistan社护罗萨他那
	f.社护罗萨他那Zabulistan.SetParent(f)

}
