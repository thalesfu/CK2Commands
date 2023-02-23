package esthonia

import (
	"github.com/thalesfu/CK2Commands/earth/scandinavia/estonia/esthonia/livs"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/estonia/esthonia/narva"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/estonia/esthonia/osel"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/estonia/esthonia/reval"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type EsthoniaDuke interface {
	feud.Duke
	CLivs利沃尼亚() livs.LivsCounty
	CNarva纳尔瓦() narva.NarvaCounty
	COsel厄瑟尔() osel.OselCounty
	CReval卡莱万() reval.RevalCounty
}

type 爱沙尼亚EsthoniaDuke struct {
	feud.BaseDuke
	利沃尼亚Livs livs.LivsCounty
	纳尔瓦Narva narva.NarvaCounty
	厄瑟尔Osel  osel.OselCounty
	卡莱万Reval reval.RevalCounty
}

func (d *爱沙尼亚EsthoniaDuke) CLivs利沃尼亚() livs.LivsCounty {
	return d.利沃尼亚Livs
}

func (d *爱沙尼亚EsthoniaDuke) CNarva纳尔瓦() narva.NarvaCounty {
	return d.纳尔瓦Narva
}

func (d *爱沙尼亚EsthoniaDuke) COsel厄瑟尔() osel.OselCounty {
	return d.厄瑟尔Osel
}

func (d *爱沙尼亚EsthoniaDuke) CReval卡莱万() reval.RevalCounty {
	return d.卡莱万Reval
}

var DEsthonia爱沙尼亚 EsthoniaDuke = &爱沙尼亚EsthoniaDuke{}

func init() {
	f := DEsthonia爱沙尼亚.(*爱沙尼亚EsthoniaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "esthonia",
		TitleName: "爱沙尼亚",
		TitleCode: "d_esthonia",
		Counties:  map[string]feud.County{},
	}

	f.利沃尼亚Livs = livs.CLivs利沃尼亚
	f.利沃尼亚Livs.SetParent(f)

	f.纳尔瓦Narva = narva.CNarva纳尔瓦
	f.纳尔瓦Narva.SetParent(f)

	f.厄瑟尔Osel = osel.COsel厄瑟尔
	f.厄瑟尔Osel.SetParent(f)

	f.卡莱万Reval = reval.CReval卡莱万
	f.卡莱万Reval.SetParent(f)

}
