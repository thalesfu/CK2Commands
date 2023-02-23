package gao

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GaoCounty interface {
	feud.County
	BAnchawadi昂沙瓦迪() feud.Barony
	BBourem布雷姆() feud.Barony
	BGabero加贝罗() feud.Barony
	BGao加奥() feud.Barony
	BGaosaney加奥_撒内() feud.Barony
	BMenaka梅纳卡() feud.Barony
	BSarnah萨尔纳() feud.Barony
}

type 加奥GaoCounty struct {
	feud.BaseCounty
	昂沙瓦迪Anchawadi feud.Barony
	布雷姆Bourem     feud.Barony
	加贝罗Gabero     feud.Barony
	加奥Gao         feud.Barony
	加奥_撒内Gaosaney feud.Barony
	梅纳卡Menaka     feud.Barony
	萨尔纳Sarnah     feud.Barony
}

func (c *加奥GaoCounty) BAnchawadi昂沙瓦迪() feud.Barony {
	return c.昂沙瓦迪Anchawadi
}

func (c *加奥GaoCounty) BBourem布雷姆() feud.Barony {
	return c.布雷姆Bourem
}

func (c *加奥GaoCounty) BGabero加贝罗() feud.Barony {
	return c.加贝罗Gabero
}

func (c *加奥GaoCounty) BGao加奥() feud.Barony {
	return c.加奥Gao
}

func (c *加奥GaoCounty) BGaosaney加奥_撒内() feud.Barony {
	return c.加奥_撒内Gaosaney
}

func (c *加奥GaoCounty) BMenaka梅纳卡() feud.Barony {
	return c.梅纳卡Menaka
}

func (c *加奥GaoCounty) BSarnah萨尔纳() feud.Barony {
	return c.萨尔纳Sarnah
}

var CGao加奥 GaoCounty = &加奥GaoCounty{}

func init() {
	f := CGao加奥.(*加奥GaoCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "914",
		Title:     "gao",
		TitleName: "加奥",
		TitleCode: "c_gao",
		Baronies:  map[string]feud.Barony{},
	}

	f.昂沙瓦迪Anchawadi = BAnchawadi昂沙瓦迪
	f.昂沙瓦迪Anchawadi.SetParent(f)

	f.布雷姆Bourem = BBourem布雷姆
	f.布雷姆Bourem.SetParent(f)

	f.加贝罗Gabero = BGabero加贝罗
	f.加贝罗Gabero.SetParent(f)

	f.加奥Gao = BGao加奥
	f.加奥Gao.SetParent(f)

	f.加奥_撒内Gaosaney = BGaosaney加奥_撒内
	f.加奥_撒内Gaosaney.SetParent(f)

	f.梅纳卡Menaka = BMenaka梅纳卡
	f.梅纳卡Menaka.SetParent(f)

	f.萨尔纳Sarnah = BSarnah萨尔纳
	f.萨尔纳Sarnah.SetParent(f)

}
