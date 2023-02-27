package al_djazair

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Al_djazairCounty interface {
    feud.County
    BAlgiers阿尔及尔() 	feud.Barony
    BChiffa希法() 	feud.Barony
    BElachour阿舒尔() 	feud.Barony
    BMirabeau德拉本海达() 	feud.Barony
    BRhedadoua里赫达多() 	feud.Barony
    BSidiakacha西迪阿卡沙() 	feud.Barony
    BTipasa提帕萨() 	feud.Barony
    BTirourda梯罗达() 	feud.Barony
    BTiziouzou蒂齐乌祖() 	feud.Barony
}

type 贾扎伊尔Al_djazairCounty struct {
	feud.BaseCounty
	阿尔及尔Algiers 	feud.Barony
	希法Chiffa 	feud.Barony
	阿舒尔Elachour 	feud.Barony
	德拉本海达Mirabeau 	feud.Barony
	里赫达多Rhedadoua 	feud.Barony
	西迪阿卡沙Sidiakacha 	feud.Barony
	提帕萨Tipasa 	feud.Barony
	梯罗达Tirourda 	feud.Barony
	蒂齐乌祖Tiziouzou 	feud.Barony
}

func (c *贾扎伊尔Al_djazairCounty) BAlgiers阿尔及尔() feud.Barony {
	return c.阿尔及尔Algiers
}
    
func (c *贾扎伊尔Al_djazairCounty) BChiffa希法() feud.Barony {
	return c.希法Chiffa
}
    
func (c *贾扎伊尔Al_djazairCounty) BElachour阿舒尔() feud.Barony {
	return c.阿舒尔Elachour
}
    
func (c *贾扎伊尔Al_djazairCounty) BMirabeau德拉本海达() feud.Barony {
	return c.德拉本海达Mirabeau
}
    
func (c *贾扎伊尔Al_djazairCounty) BRhedadoua里赫达多() feud.Barony {
	return c.里赫达多Rhedadoua
}
    
func (c *贾扎伊尔Al_djazairCounty) BSidiakacha西迪阿卡沙() feud.Barony {
	return c.西迪阿卡沙Sidiakacha
}
    
func (c *贾扎伊尔Al_djazairCounty) BTipasa提帕萨() feud.Barony {
	return c.提帕萨Tipasa
}
    
func (c *贾扎伊尔Al_djazairCounty) BTirourda梯罗达() feud.Barony {
	return c.梯罗达Tirourda
}
    
func (c *贾扎伊尔Al_djazairCounty) BTiziouzou蒂齐乌祖() feud.Barony {
	return c.蒂齐乌祖Tiziouzou
}
    
var CAl_djazair贾扎伊尔 Al_djazairCounty = &贾扎伊尔Al_djazairCounty{}

func init() {
	f := CAl_djazair贾扎伊尔.(*贾扎伊尔Al_djazairCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "831",
		Title:     "al_djazair",
		TitleName: "贾扎伊尔",
		TitleCode: "c_al_djazair",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔及尔Algiers = BAlgiers阿尔及尔
	f.阿尔及尔Algiers.SetParent(f)
	
	f.希法Chiffa = BChiffa希法
	f.希法Chiffa.SetParent(f)
	
	f.阿舒尔Elachour = BElachour阿舒尔
	f.阿舒尔Elachour.SetParent(f)
	
	f.德拉本海达Mirabeau = BMirabeau德拉本海达
	f.德拉本海达Mirabeau.SetParent(f)
	
	f.里赫达多Rhedadoua = BRhedadoua里赫达多
	f.里赫达多Rhedadoua.SetParent(f)
	
	f.西迪阿卡沙Sidiakacha = BSidiakacha西迪阿卡沙
	f.西迪阿卡沙Sidiakacha.SetParent(f)
	
	f.提帕萨Tipasa = BTipasa提帕萨
	f.提帕萨Tipasa.SetParent(f)
	
	f.梯罗达Tirourda = BTirourda梯罗达
	f.梯罗达Tirourda.SetParent(f)
	
	f.蒂齐乌祖Tiziouzou = BTiziouzou蒂齐乌祖
	f.蒂齐乌祖Tiziouzou.SetParent(f)
	
}
