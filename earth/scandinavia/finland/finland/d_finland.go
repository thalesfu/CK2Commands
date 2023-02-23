package finland

import (
	"github.com/thalesfu/CK2Commands/earth/scandinavia/finland/finland/finland"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/finland/finland/nyland"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/finland/finland/satakunta"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/finland/finland/siuntio"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/finland/finland/tavasts"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type FinlandDuke interface {
	feud.Duke
	CFinland索米() finland.FinlandCounty
	CNyland新地() nyland.NylandCounty
	CSatakunta萨塔昆塔() satakunta.SatakuntaCounty
	CSiuntio锡温蒂奥() siuntio.SiuntioCounty
	CTavasts塔瓦斯蒂亚() tavasts.TavastsCounty
}

type 芬兰FinlandDuke struct {
	feud.BaseDuke
	索米Finland     finland.FinlandCounty
	新地Nyland      nyland.NylandCounty
	萨塔昆塔Satakunta satakunta.SatakuntaCounty
	锡温蒂奥Siuntio   siuntio.SiuntioCounty
	塔瓦斯蒂亚Tavasts  tavasts.TavastsCounty
}

func (d *芬兰FinlandDuke) CFinland索米() finland.FinlandCounty {
	return d.索米Finland
}

func (d *芬兰FinlandDuke) CNyland新地() nyland.NylandCounty {
	return d.新地Nyland
}

func (d *芬兰FinlandDuke) CSatakunta萨塔昆塔() satakunta.SatakuntaCounty {
	return d.萨塔昆塔Satakunta
}

func (d *芬兰FinlandDuke) CSiuntio锡温蒂奥() siuntio.SiuntioCounty {
	return d.锡温蒂奥Siuntio
}

func (d *芬兰FinlandDuke) CTavasts塔瓦斯蒂亚() tavasts.TavastsCounty {
	return d.塔瓦斯蒂亚Tavasts
}

var DFinland芬兰 FinlandDuke = &芬兰FinlandDuke{}

func init() {
	f := DFinland芬兰.(*芬兰FinlandDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "finland",
		TitleName: "芬兰",
		TitleCode: "d_finland",
		Counties:  map[string]feud.County{},
	}

	f.索米Finland = finland.CFinland索米
	f.索米Finland.SetParent(f)

	f.新地Nyland = nyland.CNyland新地
	f.新地Nyland.SetParent(f)

	f.萨塔昆塔Satakunta = satakunta.CSatakunta萨塔昆塔
	f.萨塔昆塔Satakunta.SetParent(f)

	f.锡温蒂奥Siuntio = siuntio.CSiuntio锡温蒂奥
	f.锡温蒂奥Siuntio.SetParent(f)

	f.塔瓦斯蒂亚Tavasts = tavasts.CTavasts塔瓦斯蒂亚
	f.塔瓦斯蒂亚Tavasts.SetParent(f)

}
