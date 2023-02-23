package armenia

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/armenia/armenia/ani"
	"github.com/thalesfu/CK2Commands/earth/byzantium/armenia/armenia/dwin"
	"github.com/thalesfu/CK2Commands/earth/byzantium/armenia/armenia/lori"
	"github.com/thalesfu/CK2Commands/earth/byzantium/armenia/armenia/vaspurakan"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ArmeniaDuke interface {
	feud.Duke
	CAni阿尼() ani.AniCounty
	CDwin德温() dwin.DwinCounty
	CLori洛里() lori.LoriCounty
	CVaspurakan瓦斯普拉坎() vaspurakan.VaspurakanCounty
}

type 亚美尼亚ArmeniaDuke struct {
	feud.BaseDuke
	阿尼Ani           ani.AniCounty
	德温Dwin          dwin.DwinCounty
	洛里Lori          lori.LoriCounty
	瓦斯普拉坎Vaspurakan vaspurakan.VaspurakanCounty
}

func (d *亚美尼亚ArmeniaDuke) CAni阿尼() ani.AniCounty {
	return d.阿尼Ani
}

func (d *亚美尼亚ArmeniaDuke) CDwin德温() dwin.DwinCounty {
	return d.德温Dwin
}

func (d *亚美尼亚ArmeniaDuke) CLori洛里() lori.LoriCounty {
	return d.洛里Lori
}

func (d *亚美尼亚ArmeniaDuke) CVaspurakan瓦斯普拉坎() vaspurakan.VaspurakanCounty {
	return d.瓦斯普拉坎Vaspurakan
}

var DArmenia亚美尼亚 ArmeniaDuke = &亚美尼亚ArmeniaDuke{}

func init() {
	f := DArmenia亚美尼亚.(*亚美尼亚ArmeniaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "armenia",
		TitleName: "亚美尼亚",
		TitleCode: "d_armenia",
		Counties:  map[string]feud.County{},
	}

	f.阿尼Ani = ani.CAni阿尼
	f.阿尼Ani.SetParent(f)

	f.德温Dwin = dwin.CDwin德温
	f.德温Dwin.SetParent(f)

	f.洛里Lori = lori.CLori洛里
	f.洛里Lori.SetParent(f)

	f.瓦斯普拉坎Vaspurakan = vaspurakan.CVaspurakan瓦斯普拉坎
	f.瓦斯普拉坎Vaspurakan.SetParent(f)

}
