package taizz

import (
	"github.com/thalesfu/CK2Commands/earth/arabia/yemen/taizz/aden"
	"github.com/thalesfu/CK2Commands/earth/arabia/yemen/taizz/bayda"
	"github.com/thalesfu/CK2Commands/earth/arabia/yemen/taizz/taizz"
	"github.com/thalesfu/CK2Commands/earth/arabia/yemen/taizz/zabid"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TaizzDuke interface {
    feud.Duke
    CAden亚丁() 	aden.AdenCounty
    CBayda贝达() 	bayda.BaydaCounty
    CTaizz塔伊兹() 	taizz.TaizzCounty
    CZabid宰比德() 	zabid.ZabidCounty
}

type 塔伊兹TaizzDuke struct {
	feud.BaseDuke
	亚丁Aden 	aden.AdenCounty
	贝达Bayda 	bayda.BaydaCounty
	塔伊兹Taizz 	taizz.TaizzCounty
	宰比德Zabid 	zabid.ZabidCounty
}

func (d *塔伊兹TaizzDuke) CAden亚丁() aden.AdenCounty {
	return d.亚丁Aden
}
    
func (d *塔伊兹TaizzDuke) CBayda贝达() bayda.BaydaCounty {
	return d.贝达Bayda
}
    
func (d *塔伊兹TaizzDuke) CTaizz塔伊兹() taizz.TaizzCounty {
	return d.塔伊兹Taizz
}
    
func (d *塔伊兹TaizzDuke) CZabid宰比德() zabid.ZabidCounty {
	return d.宰比德Zabid
}
    
var DTaizz塔伊兹 TaizzDuke = &塔伊兹TaizzDuke{}

func init() {
	f := DTaizz塔伊兹.(*塔伊兹TaizzDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "taizz",
		TitleName: "塔伊兹",
		TitleCode: "d_taizz",
		Counties:  map[string]feud.County{},
	}

	f.亚丁Aden = aden.CAden亚丁
	f.亚丁Aden.SetParent(f)
	
	f.贝达Bayda = bayda.CBayda贝达
	f.贝达Bayda.SetParent(f)
	
	f.塔伊兹Taizz = taizz.CTaizz塔伊兹
	f.塔伊兹Taizz.SetParent(f)
	
	f.宰比德Zabid = zabid.CZabid宰比德
	f.宰比德Zabid.SetParent(f)
	
}
