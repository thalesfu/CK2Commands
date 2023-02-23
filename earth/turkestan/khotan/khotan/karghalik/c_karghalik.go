package karghalik

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KarghalikCounty interface {
	feud.County
	BChishi赤石() feud.Barony
	BDingkengkou丁坑口() feud.Barony
	BKarghalik喀格勒克() feud.Barony
	BQiepantuo朅盘陀() feud.Barony
	BXiachen夏臣() feud.Barony
	BZhenzigou榛子沟() feud.Barony
	BZhongqing仲庆() feud.Barony
}

type 喀格勒克KarghalikCounty struct {
	feud.BaseCounty
	赤石Chishi       feud.Barony
	丁坑口Dingkengkou feud.Barony
	喀格勒克Karghalik  feud.Barony
	朅盘陀Qiepantuo   feud.Barony
	夏臣Xiachen      feud.Barony
	榛子沟Zhenzigou   feud.Barony
	仲庆Zhongqing    feud.Barony
}

func (c *喀格勒克KarghalikCounty) BChishi赤石() feud.Barony {
	return c.赤石Chishi
}

func (c *喀格勒克KarghalikCounty) BDingkengkou丁坑口() feud.Barony {
	return c.丁坑口Dingkengkou
}

func (c *喀格勒克KarghalikCounty) BKarghalik喀格勒克() feud.Barony {
	return c.喀格勒克Karghalik
}

func (c *喀格勒克KarghalikCounty) BQiepantuo朅盘陀() feud.Barony {
	return c.朅盘陀Qiepantuo
}

func (c *喀格勒克KarghalikCounty) BXiachen夏臣() feud.Barony {
	return c.夏臣Xiachen
}

func (c *喀格勒克KarghalikCounty) BZhenzigou榛子沟() feud.Barony {
	return c.榛子沟Zhenzigou
}

func (c *喀格勒克KarghalikCounty) BZhongqing仲庆() feud.Barony {
	return c.仲庆Zhongqing
}

var CKarghalik喀格勒克 KarghalikCounty = &喀格勒克KarghalikCounty{}

func init() {
	f := CKarghalik喀格勒克.(*喀格勒克KarghalikCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1524",
		Title:     "karghalik",
		TitleName: "喀格勒克",
		TitleCode: "c_karghalik",
		Baronies:  map[string]feud.Barony{},
	}

	f.赤石Chishi = BChishi赤石
	f.赤石Chishi.SetParent(f)

	f.丁坑口Dingkengkou = BDingkengkou丁坑口
	f.丁坑口Dingkengkou.SetParent(f)

	f.喀格勒克Karghalik = BKarghalik喀格勒克
	f.喀格勒克Karghalik.SetParent(f)

	f.朅盘陀Qiepantuo = BQiepantuo朅盘陀
	f.朅盘陀Qiepantuo.SetParent(f)

	f.夏臣Xiachen = BXiachen夏臣
	f.夏臣Xiachen.SetParent(f)

	f.榛子沟Zhenzigou = BZhenzigou榛子沟
	f.榛子沟Zhenzigou.SetParent(f)

	f.仲庆Zhongqing = BZhongqing仲庆
	f.仲庆Zhongqing.SetParent(f)

}
