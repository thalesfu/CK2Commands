package don_portage

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Don_portageCounty interface {
    feud.County
    BDonskoy顿斯科伊() 	feud.Barony
    BIllevka伊利耶夫卡() 	feud.Barony
    BIllovlya伊洛夫利亚() 	feud.Barony
    BKalach卡拉奇() 	feud.Barony
    BLoq洛克() 	feud.Barony
    BOzerki奥泽尔基() 	feud.Barony
    BRyumino留米诺() 	feud.Barony
    BTary塔雷() 	feud.Barony
}

type 顿河水冲Don_portageCounty struct {
	feud.BaseCounty
	顿斯科伊Donskoy 	feud.Barony
	伊利耶夫卡Illevka 	feud.Barony
	伊洛夫利亚Illovlya 	feud.Barony
	卡拉奇Kalach 	feud.Barony
	洛克Loq 	feud.Barony
	奥泽尔基Ozerki 	feud.Barony
	留米诺Ryumino 	feud.Barony
	塔雷Tary 	feud.Barony
}

func (c *顿河水冲Don_portageCounty) BDonskoy顿斯科伊() feud.Barony {
	return c.顿斯科伊Donskoy
}
    
func (c *顿河水冲Don_portageCounty) BIllevka伊利耶夫卡() feud.Barony {
	return c.伊利耶夫卡Illevka
}
    
func (c *顿河水冲Don_portageCounty) BIllovlya伊洛夫利亚() feud.Barony {
	return c.伊洛夫利亚Illovlya
}
    
func (c *顿河水冲Don_portageCounty) BKalach卡拉奇() feud.Barony {
	return c.卡拉奇Kalach
}
    
func (c *顿河水冲Don_portageCounty) BLoq洛克() feud.Barony {
	return c.洛克Loq
}
    
func (c *顿河水冲Don_portageCounty) BOzerki奥泽尔基() feud.Barony {
	return c.奥泽尔基Ozerki
}
    
func (c *顿河水冲Don_portageCounty) BRyumino留米诺() feud.Barony {
	return c.留米诺Ryumino
}
    
func (c *顿河水冲Don_portageCounty) BTary塔雷() feud.Barony {
	return c.塔雷Tary
}
    
var CDon_portage顿河水冲 Don_portageCounty = &顿河水冲Don_portageCounty{}

func init() {
	f := CDon_portage顿河水冲.(*顿河水冲Don_portageCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "595",
		Title:     "don_portage",
		TitleName: "顿河水冲",
		TitleCode: "c_don_portage",
		Baronies:  map[string]feud.Barony{},
	}

	f.顿斯科伊Donskoy = BDonskoy顿斯科伊
	f.顿斯科伊Donskoy.SetParent(f)
	
	f.伊利耶夫卡Illevka = BIllevka伊利耶夫卡
	f.伊利耶夫卡Illevka.SetParent(f)
	
	f.伊洛夫利亚Illovlya = BIllovlya伊洛夫利亚
	f.伊洛夫利亚Illovlya.SetParent(f)
	
	f.卡拉奇Kalach = BKalach卡拉奇
	f.卡拉奇Kalach.SetParent(f)
	
	f.洛克Loq = BLoq洛克
	f.洛克Loq.SetParent(f)
	
	f.奥泽尔基Ozerki = BOzerki奥泽尔基
	f.奥泽尔基Ozerki.SetParent(f)
	
	f.留米诺Ryumino = BRyumino留米诺
	f.留米诺Ryumino.SetParent(f)
	
	f.塔雷Tary = BTary塔雷
	f.塔雷Tary.SetParent(f)
	
}
