package modena

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ModenaCounty interface {
    feud.County
    BBomporto邦波尔托() 	feud.Barony
    BCarpi卡尔皮() 	feud.Barony
    BMirandola米兰多拉() 	feud.Barony
    BModena摩德纳() 	feud.Barony
    BNovellara诺韦拉拉() 	feud.Barony
    BReggionellemila雷焦埃米利亚() 	feud.Barony
    BSabbioneta萨比奥内塔() 	feud.Barony
    BSassuolo萨索洛() 	feud.Barony
}

type 摩德纳ModenaCounty struct {
	feud.BaseCounty
	邦波尔托Bomporto 	feud.Barony
	卡尔皮Carpi 	feud.Barony
	米兰多拉Mirandola 	feud.Barony
	摩德纳Modena 	feud.Barony
	诺韦拉拉Novellara 	feud.Barony
	雷焦埃米利亚Reggionellemila 	feud.Barony
	萨比奥内塔Sabbioneta 	feud.Barony
	萨索洛Sassuolo 	feud.Barony
}

func (c *摩德纳ModenaCounty) BBomporto邦波尔托() feud.Barony {
	return c.邦波尔托Bomporto
}
    
func (c *摩德纳ModenaCounty) BCarpi卡尔皮() feud.Barony {
	return c.卡尔皮Carpi
}
    
func (c *摩德纳ModenaCounty) BMirandola米兰多拉() feud.Barony {
	return c.米兰多拉Mirandola
}
    
func (c *摩德纳ModenaCounty) BModena摩德纳() feud.Barony {
	return c.摩德纳Modena
}
    
func (c *摩德纳ModenaCounty) BNovellara诺韦拉拉() feud.Barony {
	return c.诺韦拉拉Novellara
}
    
func (c *摩德纳ModenaCounty) BReggionellemila雷焦埃米利亚() feud.Barony {
	return c.雷焦埃米利亚Reggionellemila
}
    
func (c *摩德纳ModenaCounty) BSabbioneta萨比奥内塔() feud.Barony {
	return c.萨比奥内塔Sabbioneta
}
    
func (c *摩德纳ModenaCounty) BSassuolo萨索洛() feud.Barony {
	return c.萨索洛Sassuolo
}
    
var CModena摩德纳 ModenaCounty = &摩德纳ModenaCounty{}

func init() {
	f := CModena摩德纳.(*摩德纳ModenaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "322",
		Title:     "modena",
		TitleName: "摩德纳",
		TitleCode: "c_modena",
		Baronies:  map[string]feud.Barony{},
	}

	f.邦波尔托Bomporto = BBomporto邦波尔托
	f.邦波尔托Bomporto.SetParent(f)
	
	f.卡尔皮Carpi = BCarpi卡尔皮
	f.卡尔皮Carpi.SetParent(f)
	
	f.米兰多拉Mirandola = BMirandola米兰多拉
	f.米兰多拉Mirandola.SetParent(f)
	
	f.摩德纳Modena = BModena摩德纳
	f.摩德纳Modena.SetParent(f)
	
	f.诺韦拉拉Novellara = BNovellara诺韦拉拉
	f.诺韦拉拉Novellara.SetParent(f)
	
	f.雷焦埃米利亚Reggionellemila = BReggionellemila雷焦埃米利亚
	f.雷焦埃米利亚Reggionellemila.SetParent(f)
	
	f.萨比奥内塔Sabbioneta = BSabbioneta萨比奥内塔
	f.萨比奥内塔Sabbioneta.SetParent(f)
	
	f.萨索洛Sassuolo = BSassuolo萨索洛
	f.萨索洛Sassuolo.SetParent(f)
	
}
