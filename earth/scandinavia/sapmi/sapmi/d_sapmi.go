package sapmi

import (
	"github.com/thalesfu/CK2Commands/earth/scandinavia/sapmi/sapmi/lappland"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/sapmi/sapmi/vasterbotten"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SapmiDuke interface {
	feud.Duke
	CLappland拉普兰() lappland.LapplandCounty
	CVasterbotten西博滕() vasterbotten.VasterbottenCounty
}

type 拉普兰SapmiDuke struct {
	feud.BaseDuke
	拉普兰Lappland     lappland.LapplandCounty
	西博滕Vasterbotten vasterbotten.VasterbottenCounty
}

func (d *拉普兰SapmiDuke) CLappland拉普兰() lappland.LapplandCounty {
	return d.拉普兰Lappland
}

func (d *拉普兰SapmiDuke) CVasterbotten西博滕() vasterbotten.VasterbottenCounty {
	return d.西博滕Vasterbotten
}

var DSapmi拉普兰 SapmiDuke = &拉普兰SapmiDuke{}

func init() {
	f := DSapmi拉普兰.(*拉普兰SapmiDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "sapmi",
		TitleName: "拉普兰",
		TitleCode: "d_sapmi",
		Counties:  map[string]feud.County{},
	}

	f.拉普兰Lappland = lappland.CLappland拉普兰
	f.拉普兰Lappland.SetParent(f)

	f.西博滕Vasterbotten = vasterbotten.CVasterbotten西博滕
	f.西博滕Vasterbotten.SetParent(f)

}
