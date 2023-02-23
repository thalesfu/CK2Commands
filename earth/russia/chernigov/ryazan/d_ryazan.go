package ryazan

import (
	"github.com/thalesfu/CK2Commands/earth/russia/chernigov/ryazan/kolomna"
	"github.com/thalesfu/CK2Commands/earth/russia/chernigov/ryazan/pronsk"
	"github.com/thalesfu/CK2Commands/earth/russia/chernigov/ryazan/ryazan"
	"github.com/thalesfu/CK2Commands/earth/russia/chernigov/ryazan/yelets"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RyazanDuke interface {
	feud.Duke
	CKolomna图拉() kolomna.KolomnaCounty
	CPronsk普龙斯克() pronsk.PronskCounty
	CRyazan梁赞() ryazan.RyazanCounty
	CYelets叶列茨() yelets.YeletsCounty
}

type 梁赞RyazanDuke struct {
	feud.BaseDuke
	图拉Kolomna  kolomna.KolomnaCounty
	普龙斯克Pronsk pronsk.PronskCounty
	梁赞Ryazan   ryazan.RyazanCounty
	叶列茨Yelets  yelets.YeletsCounty
}

func (d *梁赞RyazanDuke) CKolomna图拉() kolomna.KolomnaCounty {
	return d.图拉Kolomna
}

func (d *梁赞RyazanDuke) CPronsk普龙斯克() pronsk.PronskCounty {
	return d.普龙斯克Pronsk
}

func (d *梁赞RyazanDuke) CRyazan梁赞() ryazan.RyazanCounty {
	return d.梁赞Ryazan
}

func (d *梁赞RyazanDuke) CYelets叶列茨() yelets.YeletsCounty {
	return d.叶列茨Yelets
}

var DRyazan梁赞 RyazanDuke = &梁赞RyazanDuke{}

func init() {
	f := DRyazan梁赞.(*梁赞RyazanDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "ryazan",
		TitleName: "梁赞",
		TitleCode: "d_ryazan",
		Counties:  map[string]feud.County{},
	}

	f.图拉Kolomna = kolomna.CKolomna图拉
	f.图拉Kolomna.SetParent(f)

	f.普龙斯克Pronsk = pronsk.CPronsk普龙斯克
	f.普龙斯克Pronsk.SetParent(f)

	f.梁赞Ryazan = ryazan.CRyazan梁赞
	f.梁赞Ryazan.SetParent(f)

	f.叶列茨Yelets = yelets.CYelets叶列茨
	f.叶列茨Yelets.SetParent(f)

}
