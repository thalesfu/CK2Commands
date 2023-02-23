package lombardia

import (
	"github.com/thalesfu/CK2Commands/earth/italy/italy/lombardia/brescia"
	"github.com/thalesfu/CK2Commands/earth/italy/italy/lombardia/como"
	"github.com/thalesfu/CK2Commands/earth/italy/italy/lombardia/cremona"
	"github.com/thalesfu/CK2Commands/earth/italy/italy/lombardia/leventina"
	"github.com/thalesfu/CK2Commands/earth/italy/italy/lombardia/lombardia"
	"github.com/thalesfu/CK2Commands/earth/italy/italy/lombardia/pavia"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LombardiaDuke interface {
	feud.Duke
	CBrescia布雷西亚() brescia.BresciaCounty
	CComo科莫() como.ComoCounty
	CCremona克雷莫纳() cremona.CremonaCounty
	CLeventina贝林佐纳() leventina.LeventinaCounty
	CLombardia米兰() lombardia.LombardiaCounty
	CPavia帕维亚() pavia.PaviaCounty
}

type 米兰LombardiaDuke struct {
	feud.BaseDuke
	布雷西亚Brescia   brescia.BresciaCounty
	科莫Como        como.ComoCounty
	克雷莫纳Cremona   cremona.CremonaCounty
	贝林佐纳Leventina leventina.LeventinaCounty
	米兰Lombardia   lombardia.LombardiaCounty
	帕维亚Pavia      pavia.PaviaCounty
}

func (d *米兰LombardiaDuke) CBrescia布雷西亚() brescia.BresciaCounty {
	return d.布雷西亚Brescia
}

func (d *米兰LombardiaDuke) CComo科莫() como.ComoCounty {
	return d.科莫Como
}

func (d *米兰LombardiaDuke) CCremona克雷莫纳() cremona.CremonaCounty {
	return d.克雷莫纳Cremona
}

func (d *米兰LombardiaDuke) CLeventina贝林佐纳() leventina.LeventinaCounty {
	return d.贝林佐纳Leventina
}

func (d *米兰LombardiaDuke) CLombardia米兰() lombardia.LombardiaCounty {
	return d.米兰Lombardia
}

func (d *米兰LombardiaDuke) CPavia帕维亚() pavia.PaviaCounty {
	return d.帕维亚Pavia
}

var DLombardia米兰 LombardiaDuke = &米兰LombardiaDuke{}

func init() {
	f := DLombardia米兰.(*米兰LombardiaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "lombardia",
		TitleName: "米兰",
		TitleCode: "d_lombardia",
		Counties:  map[string]feud.County{},
	}

	f.布雷西亚Brescia = brescia.CBrescia布雷西亚
	f.布雷西亚Brescia.SetParent(f)

	f.科莫Como = como.CComo科莫
	f.科莫Como.SetParent(f)

	f.克雷莫纳Cremona = cremona.CCremona克雷莫纳
	f.克雷莫纳Cremona.SetParent(f)

	f.贝林佐纳Leventina = leventina.CLeventina贝林佐纳
	f.贝林佐纳Leventina.SetParent(f)

	f.米兰Lombardia = lombardia.CLombardia米兰
	f.米兰Lombardia.SetParent(f)

	f.帕维亚Pavia = pavia.CPavia帕维亚
	f.帕维亚Pavia.SetParent(f)

}
