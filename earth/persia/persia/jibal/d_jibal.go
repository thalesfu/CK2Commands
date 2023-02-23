package jibal

import (
	"github.com/thalesfu/CK2Commands/earth/persia/persia/jibal/esfahan"
	"github.com/thalesfu/CK2Commands/earth/persia/persia/jibal/hamadan"
	"github.com/thalesfu/CK2Commands/earth/persia/persia/jibal/luristan"
	"github.com/thalesfu/CK2Commands/earth/persia/persia/jibal/qom"
	"github.com/thalesfu/CK2Commands/earth/persia/persia/jibal/rayy"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type JibalDuke interface {
	feud.Duke
	CEsfahan伊斯法罕() esfahan.EsfahanCounty
	CHamadan哈马丹() hamadan.HamadanCounty
	CLuristan洛雷斯坦() luristan.LuristanCounty
	CQom库姆() qom.QomCounty
	CRayy雷伊() rayy.RayyCounty
}

type 吉巴勒JibalDuke struct {
	feud.BaseDuke
	伊斯法罕Esfahan  esfahan.EsfahanCounty
	哈马丹Hamadan   hamadan.HamadanCounty
	洛雷斯坦Luristan luristan.LuristanCounty
	库姆Qom        qom.QomCounty
	雷伊Rayy       rayy.RayyCounty
}

func (d *吉巴勒JibalDuke) CEsfahan伊斯法罕() esfahan.EsfahanCounty {
	return d.伊斯法罕Esfahan
}

func (d *吉巴勒JibalDuke) CHamadan哈马丹() hamadan.HamadanCounty {
	return d.哈马丹Hamadan
}

func (d *吉巴勒JibalDuke) CLuristan洛雷斯坦() luristan.LuristanCounty {
	return d.洛雷斯坦Luristan
}

func (d *吉巴勒JibalDuke) CQom库姆() qom.QomCounty {
	return d.库姆Qom
}

func (d *吉巴勒JibalDuke) CRayy雷伊() rayy.RayyCounty {
	return d.雷伊Rayy
}

var DJibal吉巴勒 JibalDuke = &吉巴勒JibalDuke{}

func init() {
	f := DJibal吉巴勒.(*吉巴勒JibalDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "jibal",
		TitleName: "吉巴勒",
		TitleCode: "d_jibal",
		Counties:  map[string]feud.County{},
	}

	f.伊斯法罕Esfahan = esfahan.CEsfahan伊斯法罕
	f.伊斯法罕Esfahan.SetParent(f)

	f.哈马丹Hamadan = hamadan.CHamadan哈马丹
	f.哈马丹Hamadan.SetParent(f)

	f.洛雷斯坦Luristan = luristan.CLuristan洛雷斯坦
	f.洛雷斯坦Luristan.SetParent(f)

	f.库姆Qom = qom.CQom库姆
	f.库姆Qom.SetParent(f)

	f.雷伊Rayy = rayy.CRayy雷伊
	f.雷伊Rayy.SetParent(f)

}
