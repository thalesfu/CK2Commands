package dandakaranya

import (
	"github.com/thalesfu/CK2Commands/earth/bengal/orissa/dandakaranya/barasuru"
	"github.com/thalesfu/CK2Commands/earth/bengal/orissa/dandakaranya/cakrakuta"
	"github.com/thalesfu/CK2Commands/earth/bengal/orissa/dandakaranya/nandapur"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DandakaranyaDuke interface {
	feud.Duke
	CBarasuru婆罗须卢() barasuru.BarasuruCounty
	CCakrakuta遮迦罗矩吒() cakrakuta.CakrakutaCounty
	CNandapur难陀城() nandapur.NandapurCounty
}

type 弹宅迦林DandakaranyaDuke struct {
	feud.BaseDuke
	婆罗须卢Barasuru   barasuru.BarasuruCounty
	遮迦罗矩吒Cakrakuta cakrakuta.CakrakutaCounty
	难陀城Nandapur    nandapur.NandapurCounty
}

func (d *弹宅迦林DandakaranyaDuke) CBarasuru婆罗须卢() barasuru.BarasuruCounty {
	return d.婆罗须卢Barasuru
}

func (d *弹宅迦林DandakaranyaDuke) CCakrakuta遮迦罗矩吒() cakrakuta.CakrakutaCounty {
	return d.遮迦罗矩吒Cakrakuta
}

func (d *弹宅迦林DandakaranyaDuke) CNandapur难陀城() nandapur.NandapurCounty {
	return d.难陀城Nandapur
}

var DDandakaranya弹宅迦林 DandakaranyaDuke = &弹宅迦林DandakaranyaDuke{}

func init() {
	f := DDandakaranya弹宅迦林.(*弹宅迦林DandakaranyaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "dandakaranya",
		TitleName: "弹宅迦林",
		TitleCode: "d_dandakaranya",
		Counties:  map[string]feud.County{},
	}

	f.婆罗须卢Barasuru = barasuru.CBarasuru婆罗须卢
	f.婆罗须卢Barasuru.SetParent(f)

	f.遮迦罗矩吒Cakrakuta = cakrakuta.CCakrakuta遮迦罗矩吒
	f.遮迦罗矩吒Cakrakuta.SetParent(f)

	f.难陀城Nandapur = nandapur.CNandapur难陀城
	f.难陀城Nandapur.SetParent(f)

}
