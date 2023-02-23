package tao

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/georgia/tao/guria"
	"github.com/thalesfu/CK2Commands/earth/byzantium/georgia/tao/klarjeti"
	"github.com/thalesfu/CK2Commands/earth/byzantium/georgia/tao/tao"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TaoDuke interface {
	feud.Duke
	CGuria古里亚() guria.GuriaCounty
	CKlarjeti克拉尔哲季() klarjeti.KlarjetiCounty
	CTao陶() tao.TaoCounty
}

type 陶TaoDuke struct {
	feud.BaseDuke
	古里亚Guria      guria.GuriaCounty
	克拉尔哲季Klarjeti klarjeti.KlarjetiCounty
	陶Tao          tao.TaoCounty
}

func (d *陶TaoDuke) CGuria古里亚() guria.GuriaCounty {
	return d.古里亚Guria
}

func (d *陶TaoDuke) CKlarjeti克拉尔哲季() klarjeti.KlarjetiCounty {
	return d.克拉尔哲季Klarjeti
}

func (d *陶TaoDuke) CTao陶() tao.TaoCounty {
	return d.陶Tao
}

var DTao陶 TaoDuke = &陶TaoDuke{}

func init() {
	f := DTao陶.(*陶TaoDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "tao",
		TitleName: "陶",
		TitleCode: "d_tao",
		Counties:  map[string]feud.County{},
	}

	f.古里亚Guria = guria.CGuria古里亚
	f.古里亚Guria.SetParent(f)

	f.克拉尔哲季Klarjeti = klarjeti.CKlarjeti克拉尔哲季
	f.克拉尔哲季Klarjeti.SetParent(f)

	f.陶Tao = tao.CTao陶
	f.陶Tao.SetParent(f)

}
