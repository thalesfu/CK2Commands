package toscana

import (
	"github.com/thalesfu/CK2Commands/earth/italy/italy/toscana/arezzo"
	"github.com/thalesfu/CK2Commands/earth/italy/italy/toscana/firenze"
	"github.com/thalesfu/CK2Commands/earth/italy/italy/toscana/lucca"
	"github.com/thalesfu/CK2Commands/earth/italy/italy/toscana/siena"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ToscanaDuke interface {
	feud.Duke
	CArezzo阿雷佐() arezzo.ArezzoCounty
	CFirenze佛罗伦萨() firenze.FirenzeCounty
	CLucca卢卡() lucca.LuccaCounty
	CSiena锡耶纳() siena.SienaCounty
}

type 托斯卡纳ToscanaDuke struct {
	feud.BaseDuke
	阿雷佐Arezzo   arezzo.ArezzoCounty
	佛罗伦萨Firenze firenze.FirenzeCounty
	卢卡Lucca     lucca.LuccaCounty
	锡耶纳Siena    siena.SienaCounty
}

func (d *托斯卡纳ToscanaDuke) CArezzo阿雷佐() arezzo.ArezzoCounty {
	return d.阿雷佐Arezzo
}

func (d *托斯卡纳ToscanaDuke) CFirenze佛罗伦萨() firenze.FirenzeCounty {
	return d.佛罗伦萨Firenze
}

func (d *托斯卡纳ToscanaDuke) CLucca卢卡() lucca.LuccaCounty {
	return d.卢卡Lucca
}

func (d *托斯卡纳ToscanaDuke) CSiena锡耶纳() siena.SienaCounty {
	return d.锡耶纳Siena
}

var DToscana托斯卡纳 ToscanaDuke = &托斯卡纳ToscanaDuke{}

func init() {
	f := DToscana托斯卡纳.(*托斯卡纳ToscanaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "toscana",
		TitleName: "托斯卡纳",
		TitleCode: "d_toscana",
		Counties:  map[string]feud.County{},
	}

	f.阿雷佐Arezzo = arezzo.CArezzo阿雷佐
	f.阿雷佐Arezzo.SetParent(f)

	f.佛罗伦萨Firenze = firenze.CFirenze佛罗伦萨
	f.佛罗伦萨Firenze.SetParent(f)

	f.卢卡Lucca = lucca.CLucca卢卡
	f.卢卡Lucca.SetParent(f)

	f.锡耶纳Siena = siena.CSiena锡耶纳
	f.锡耶纳Siena.SetParent(f)

}
