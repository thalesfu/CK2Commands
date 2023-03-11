package galich

import (
	"github.com/thalesfu/CK2Commands/earth/russia/galicia_volhynia/galich/galich"
	"github.com/thalesfu/CK2Commands/earth/russia/galicia_volhynia/galich/terebovl"
	"github.com/thalesfu/CK2Commands/earth/russia/galicia_volhynia/galich/ushytsia"
	"github.com/thalesfu/CK2Commands/earth/russia/galicia_volhynia/galich/zaslav"
	"github.com/thalesfu/CK2Commands/earth/russia/galicia_volhynia/galich/zvenyhorod"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GalichDuke interface {
    feud.Duke
    CGalich加利奇() 	galich.GalichCounty
    CTerebovl捷列博夫尔() 	terebovl.TerebovlCounty
    CUshytsia乌希察() 	ushytsia.UshytsiaCounty
    CZaslav扎斯拉夫() 	zaslav.ZaslavCounty
    CZvenyhorod兹韦尼戈罗德() 	zvenyhorod.ZvenyhorodCounty
}

type 加利奇GalichDuke struct {
	feud.BaseDuke
	加利奇Galich 	galich.GalichCounty
	捷列博夫尔Terebovl 	terebovl.TerebovlCounty
	乌希察Ushytsia 	ushytsia.UshytsiaCounty
	扎斯拉夫Zaslav 	zaslav.ZaslavCounty
	兹韦尼戈罗德Zvenyhorod 	zvenyhorod.ZvenyhorodCounty
}

func (d *加利奇GalichDuke) CGalich加利奇() galich.GalichCounty {
	return d.加利奇Galich
}
    
func (d *加利奇GalichDuke) CTerebovl捷列博夫尔() terebovl.TerebovlCounty {
	return d.捷列博夫尔Terebovl
}
    
func (d *加利奇GalichDuke) CUshytsia乌希察() ushytsia.UshytsiaCounty {
	return d.乌希察Ushytsia
}
    
func (d *加利奇GalichDuke) CZaslav扎斯拉夫() zaslav.ZaslavCounty {
	return d.扎斯拉夫Zaslav
}
    
func (d *加利奇GalichDuke) CZvenyhorod兹韦尼戈罗德() zvenyhorod.ZvenyhorodCounty {
	return d.兹韦尼戈罗德Zvenyhorod
}
    
var DGalich加利奇 GalichDuke = &加利奇GalichDuke{}

func init() {
	f := DGalich加利奇.(*加利奇GalichDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "galich",
		TitleName: "加利奇",
		TitleCode: "d_galich",
		Counties:  map[string]feud.County{},
	}

	f.加利奇Galich = galich.CGalich加利奇
	f.加利奇Galich.SetParent(f)
	
	f.捷列博夫尔Terebovl = terebovl.CTerebovl捷列博夫尔
	f.捷列博夫尔Terebovl.SetParent(f)
	
	f.乌希察Ushytsia = ushytsia.CUshytsia乌希察
	f.乌希察Ushytsia.SetParent(f)
	
	f.扎斯拉夫Zaslav = zaslav.CZaslav扎斯拉夫
	f.扎斯拉夫Zaslav.SetParent(f)
	
	f.兹韦尼戈罗德Zvenyhorod = zvenyhorod.CZvenyhorod兹韦尼戈罗德
	f.兹韦尼戈罗德Zvenyhorod.SetParent(f)
	
}
