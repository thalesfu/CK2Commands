package purang

import (
	"github.com/thalesfu/CK2Commands/earth/tibet/guge/purang/gegyai"
	"github.com/thalesfu/CK2Commands/earth/tibet/guge/purang/kyunglung"
	"github.com/thalesfu/CK2Commands/earth/tibet/guge/purang/purang"
	"github.com/thalesfu/CK2Commands/earth/tibet/guge/purang/tsaparang"
	"github.com/thalesfu/CK2Commands/earth/tibet/guge/purang/zhongba"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PurangDuke interface {
	feud.Duke
	CGegyai革吉() gegyai.GegyaiCounty
	CKyunglung琼隆() kyunglung.KyunglungCounty
	CPurang布让() purang.PurangCounty
	CTsaparang札布让() tsaparang.TsaparangCounty
	CZhongba仲巴() zhongba.ZhongbaCounty
}

type 布让PurangDuke struct {
	feud.BaseDuke
	革吉Gegyai     gegyai.GegyaiCounty
	琼隆Kyunglung  kyunglung.KyunglungCounty
	布让Purang     purang.PurangCounty
	札布让Tsaparang tsaparang.TsaparangCounty
	仲巴Zhongba    zhongba.ZhongbaCounty
}

func (d *布让PurangDuke) CGegyai革吉() gegyai.GegyaiCounty {
	return d.革吉Gegyai
}

func (d *布让PurangDuke) CKyunglung琼隆() kyunglung.KyunglungCounty {
	return d.琼隆Kyunglung
}

func (d *布让PurangDuke) CPurang布让() purang.PurangCounty {
	return d.布让Purang
}

func (d *布让PurangDuke) CTsaparang札布让() tsaparang.TsaparangCounty {
	return d.札布让Tsaparang
}

func (d *布让PurangDuke) CZhongba仲巴() zhongba.ZhongbaCounty {
	return d.仲巴Zhongba
}

var DPurang布让 PurangDuke = &布让PurangDuke{}

func init() {
	f := DPurang布让.(*布让PurangDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "purang",
		TitleName: "布让",
		TitleCode: "d_purang",
		Counties:  map[string]feud.County{},
	}

	f.革吉Gegyai = gegyai.CGegyai革吉
	f.革吉Gegyai.SetParent(f)

	f.琼隆Kyunglung = kyunglung.CKyunglung琼隆
	f.琼隆Kyunglung.SetParent(f)

	f.布让Purang = purang.CPurang布让
	f.布让Purang.SetParent(f)

	f.札布让Tsaparang = tsaparang.CTsaparang札布让
	f.札布让Tsaparang.SetParent(f)

	f.仲巴Zhongba = zhongba.CZhongba仲巴
	f.仲巴Zhongba.SetParent(f)

}
