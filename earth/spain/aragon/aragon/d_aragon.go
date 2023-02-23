package aragon

import (
	"github.com/thalesfu/CK2Commands/earth/spain/aragon/aragon/albarracin"
	"github.com/thalesfu/CK2Commands/earth/spain/aragon/aragon/calatayud"
	"github.com/thalesfu/CK2Commands/earth/spain/aragon/aragon/zaragoza"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AragonDuke interface {
	feud.Duke
	CAlbarracin阿尔瓦拉辛() albarracin.AlbarracinCounty
	CCalatayud卡拉塔尤() calatayud.CalatayudCounty
	CZaragoza萨拉戈萨() zaragoza.ZaragozaCounty
}

type 阿拉贡AragonDuke struct {
	feud.BaseDuke
	阿尔瓦拉辛Albarracin albarracin.AlbarracinCounty
	卡拉塔尤Calatayud   calatayud.CalatayudCounty
	萨拉戈萨Zaragoza    zaragoza.ZaragozaCounty
}

func (d *阿拉贡AragonDuke) CAlbarracin阿尔瓦拉辛() albarracin.AlbarracinCounty {
	return d.阿尔瓦拉辛Albarracin
}

func (d *阿拉贡AragonDuke) CCalatayud卡拉塔尤() calatayud.CalatayudCounty {
	return d.卡拉塔尤Calatayud
}

func (d *阿拉贡AragonDuke) CZaragoza萨拉戈萨() zaragoza.ZaragozaCounty {
	return d.萨拉戈萨Zaragoza
}

var DAragon阿拉贡 AragonDuke = &阿拉贡AragonDuke{}

func init() {
	f := DAragon阿拉贡.(*阿拉贡AragonDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "aragon",
		TitleName: "阿拉贡",
		TitleCode: "d_aragon",
		Counties:  map[string]feud.County{},
	}

	f.阿尔瓦拉辛Albarracin = albarracin.CAlbarracin阿尔瓦拉辛
	f.阿尔瓦拉辛Albarracin.SetParent(f)

	f.卡拉塔尤Calatayud = calatayud.CCalatayud卡拉塔尤
	f.卡拉塔尤Calatayud.SetParent(f)

	f.萨拉戈萨Zaragoza = zaragoza.CZaragoza萨拉戈萨
	f.萨拉戈萨Zaragoza.SetParent(f)

}
