package fars

import (
	"github.com/thalesfu/CK2Commands/earth/persia/persia/fars/fars"
	"github.com/thalesfu/CK2Commands/earth/persia/persia/fars/hendjan"
	"github.com/thalesfu/CK2Commands/earth/persia/persia/fars/istakhr"
	"github.com/thalesfu/CK2Commands/earth/persia/persia/fars/ladistan"
	"github.com/thalesfu/CK2Commands/earth/persia/persia/fars/shiraz"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type FarsDuke interface {
	feud.Duke
	CFars达什特斯坦() fars.FarsCounty
	CHendjan阿拉詹() hendjan.HendjanCounty
	CIstakhr伊斯塔赫尔() istakhr.IstakhrCounty
	CLadistan达拉布杰德() ladistan.LadistanCounty
	CShiraz设拉子() shiraz.ShirazCounty
}

type 法尔斯FarsDuke struct {
	feud.BaseDuke
	达什特斯坦Fars     fars.FarsCounty
	阿拉詹Hendjan    hendjan.HendjanCounty
	伊斯塔赫尔Istakhr  istakhr.IstakhrCounty
	达拉布杰德Ladistan ladistan.LadistanCounty
	设拉子Shiraz     shiraz.ShirazCounty
}

func (d *法尔斯FarsDuke) CFars达什特斯坦() fars.FarsCounty {
	return d.达什特斯坦Fars
}

func (d *法尔斯FarsDuke) CHendjan阿拉詹() hendjan.HendjanCounty {
	return d.阿拉詹Hendjan
}

func (d *法尔斯FarsDuke) CIstakhr伊斯塔赫尔() istakhr.IstakhrCounty {
	return d.伊斯塔赫尔Istakhr
}

func (d *法尔斯FarsDuke) CLadistan达拉布杰德() ladistan.LadistanCounty {
	return d.达拉布杰德Ladistan
}

func (d *法尔斯FarsDuke) CShiraz设拉子() shiraz.ShirazCounty {
	return d.设拉子Shiraz
}

var DFars法尔斯 FarsDuke = &法尔斯FarsDuke{}

func init() {
	f := DFars法尔斯.(*法尔斯FarsDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "fars",
		TitleName: "法尔斯",
		TitleCode: "d_fars",
		Counties:  map[string]feud.County{},
	}

	f.达什特斯坦Fars = fars.CFars达什特斯坦
	f.达什特斯坦Fars.SetParent(f)

	f.阿拉詹Hendjan = hendjan.CHendjan阿拉詹
	f.阿拉詹Hendjan.SetParent(f)

	f.伊斯塔赫尔Istakhr = istakhr.CIstakhr伊斯塔赫尔
	f.伊斯塔赫尔Istakhr.SetParent(f)

	f.达拉布杰德Ladistan = ladistan.CLadistan达拉布杰德
	f.达拉布杰德Ladistan.SetParent(f)

	f.设拉子Shiraz = shiraz.CShiraz设拉子
	f.设拉子Shiraz.SetParent(f)

}
