package jiuquan

import (
	"github.com/thalesfu/CK2Commands/earth/tibet/xixia/jiuquan/anxi"
	"github.com/thalesfu/CK2Commands/earth/tibet/xixia/jiuquan/dunhuang"
	"github.com/thalesfu/CK2Commands/earth/tibet/xixia/jiuquan/jiuquan"
	"github.com/thalesfu/CK2Commands/earth/tibet/xixia/jiuquan/yumen"
	"github.com/thalesfu/CK2Commands/earth/tibet/xixia/jiuquan/yungguan"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type JiuquanDuke interface {
	feud.Duke
	CAnxi安西() anxi.AnxiCounty
	CDunhuang敦煌() dunhuang.DunhuangCounty
	CJiuquan酒泉() jiuquan.JiuquanCounty
	CYumen玉门() yumen.YumenCounty
	CYungguan阳关() yungguan.YungguanCounty
}

type 酒泉JiuquanDuke struct {
	feud.BaseDuke
	安西Anxi     anxi.AnxiCounty
	敦煌Dunhuang dunhuang.DunhuangCounty
	酒泉Jiuquan  jiuquan.JiuquanCounty
	玉门Yumen    yumen.YumenCounty
	阳关Yungguan yungguan.YungguanCounty
}

func (d *酒泉JiuquanDuke) CAnxi安西() anxi.AnxiCounty {
	return d.安西Anxi
}

func (d *酒泉JiuquanDuke) CDunhuang敦煌() dunhuang.DunhuangCounty {
	return d.敦煌Dunhuang
}

func (d *酒泉JiuquanDuke) CJiuquan酒泉() jiuquan.JiuquanCounty {
	return d.酒泉Jiuquan
}

func (d *酒泉JiuquanDuke) CYumen玉门() yumen.YumenCounty {
	return d.玉门Yumen
}

func (d *酒泉JiuquanDuke) CYungguan阳关() yungguan.YungguanCounty {
	return d.阳关Yungguan
}

var DJiuquan酒泉 JiuquanDuke = &酒泉JiuquanDuke{}

func init() {
	f := DJiuquan酒泉.(*酒泉JiuquanDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "jiuquan",
		TitleName: "酒泉",
		TitleCode: "d_jiuquan",
		Counties:  map[string]feud.County{},
	}

	f.安西Anxi = anxi.CAnxi安西
	f.安西Anxi.SetParent(f)

	f.敦煌Dunhuang = dunhuang.CDunhuang敦煌
	f.敦煌Dunhuang.SetParent(f)

	f.酒泉Jiuquan = jiuquan.CJiuquan酒泉
	f.酒泉Jiuquan.SetParent(f)

	f.玉门Yumen = yumen.CYumen玉门
	f.玉门Yumen.SetParent(f)

	f.阳关Yungguan = yungguan.CYungguan阳关
	f.阳关Yungguan.SetParent(f)

}
