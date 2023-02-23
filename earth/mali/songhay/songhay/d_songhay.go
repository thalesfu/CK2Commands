package songhay

import (
	"github.com/thalesfu/CK2Commands/earth/mali/songhay/songhay/gao"
	"github.com/thalesfu/CK2Commands/earth/mali/songhay/songhay/kurumba"
	"github.com/thalesfu/CK2Commands/earth/mali/songhay/songhay/songhay"
	"github.com/thalesfu/CK2Commands/earth/mali/songhay/songhay/tadmekka"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SonghayDuke interface {
	feud.Duke
	CGao加奥() gao.GaoCounty
	CKurumba古伦巴() kurumba.KurumbaCounty
	CSonghay桑海() songhay.SonghayCounty
	CTadmekka塔得迈卡() tadmekka.TadmekkaCounty
}

type 桑海SonghayDuke struct {
	feud.BaseDuke
	加奥Gao        gao.GaoCounty
	古伦巴Kurumba   kurumba.KurumbaCounty
	桑海Songhay    songhay.SonghayCounty
	塔得迈卡Tadmekka tadmekka.TadmekkaCounty
}

func (d *桑海SonghayDuke) CGao加奥() gao.GaoCounty {
	return d.加奥Gao
}

func (d *桑海SonghayDuke) CKurumba古伦巴() kurumba.KurumbaCounty {
	return d.古伦巴Kurumba
}

func (d *桑海SonghayDuke) CSonghay桑海() songhay.SonghayCounty {
	return d.桑海Songhay
}

func (d *桑海SonghayDuke) CTadmekka塔得迈卡() tadmekka.TadmekkaCounty {
	return d.塔得迈卡Tadmekka
}

var DSonghay桑海 SonghayDuke = &桑海SonghayDuke{}

func init() {
	f := DSonghay桑海.(*桑海SonghayDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "songhay",
		TitleName: "桑海",
		TitleCode: "d_songhay",
		Counties:  map[string]feud.County{},
	}

	f.加奥Gao = gao.CGao加奥
	f.加奥Gao.SetParent(f)

	f.古伦巴Kurumba = kurumba.CKurumba古伦巴
	f.古伦巴Kurumba.SetParent(f)

	f.桑海Songhay = songhay.CSonghay桑海
	f.桑海Songhay.SetParent(f)

	f.塔得迈卡Tadmekka = tadmekka.CTadmekka塔得迈卡
	f.塔得迈卡Tadmekka.SetParent(f)

}
