package luntai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LuntaiCounty interface {
	feud.County
	BChangji昌吉() feud.Barony
	BDanhuan单桓() feud.Barony
	BFukang阜康() feud.Barony
	BLuntai轮台() feud.Barony
	BManas玛纳斯() feud.Barony
	BQiemi且弥() feud.Barony
	BWutanzili乌贪訾离() feud.Barony
}

type 轮台LuntaiCounty struct {
	feud.BaseCounty
	昌吉Changji     feud.Barony
	单桓Danhuan     feud.Barony
	阜康Fukang      feud.Barony
	轮台Luntai      feud.Barony
	玛纳斯Manas      feud.Barony
	且弥Qiemi       feud.Barony
	乌贪訾离Wutanzili feud.Barony
}

func (c *轮台LuntaiCounty) BChangji昌吉() feud.Barony {
	return c.昌吉Changji
}

func (c *轮台LuntaiCounty) BDanhuan单桓() feud.Barony {
	return c.单桓Danhuan
}

func (c *轮台LuntaiCounty) BFukang阜康() feud.Barony {
	return c.阜康Fukang
}

func (c *轮台LuntaiCounty) BLuntai轮台() feud.Barony {
	return c.轮台Luntai
}

func (c *轮台LuntaiCounty) BManas玛纳斯() feud.Barony {
	return c.玛纳斯Manas
}

func (c *轮台LuntaiCounty) BQiemi且弥() feud.Barony {
	return c.且弥Qiemi
}

func (c *轮台LuntaiCounty) BWutanzili乌贪訾离() feud.Barony {
	return c.乌贪訾离Wutanzili
}

var CLuntai轮台 LuntaiCounty = &轮台LuntaiCounty{}

func init() {
	f := CLuntai轮台.(*轮台LuntaiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1449",
		Title:     "luntai",
		TitleName: "轮台",
		TitleCode: "c_luntai",
		Baronies:  map[string]feud.Barony{},
	}

	f.昌吉Changji = BChangji昌吉
	f.昌吉Changji.SetParent(f)

	f.单桓Danhuan = BDanhuan单桓
	f.单桓Danhuan.SetParent(f)

	f.阜康Fukang = BFukang阜康
	f.阜康Fukang.SetParent(f)

	f.轮台Luntai = BLuntai轮台
	f.轮台Luntai.SetParent(f)

	f.玛纳斯Manas = BManas玛纳斯
	f.玛纳斯Manas.SetParent(f)

	f.且弥Qiemi = BQiemi且弥
	f.且弥Qiemi.SetParent(f)

	f.乌贪訾离Wutanzili = BWutanzili乌贪訾离
	f.乌贪訾离Wutanzili.SetParent(f)

}
