package nishapur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NishapurCounty interface {
    feud.County
    BAkhlamad阿赫拉马德() 	feud.Barony
    BChenaran恰纳兰() 	feud.Barony
    BDaregaz达尔加兹() 	feud.Barony
    BJajarm贾贾尔姆() 	feud.Barony
    BNishapur内沙布尔() 	feud.Barony
    BQuchan古昌() 	feud.Barony
    BSabzevar萨卜泽瓦尔() 	feud.Barony
    BShamat沙马特() 	feud.Barony
}

type 内沙布尔NishapurCounty struct {
	feud.BaseCounty
	阿赫拉马德Akhlamad 	feud.Barony
	恰纳兰Chenaran 	feud.Barony
	达尔加兹Daregaz 	feud.Barony
	贾贾尔姆Jajarm 	feud.Barony
	内沙布尔Nishapur 	feud.Barony
	古昌Quchan 	feud.Barony
	萨卜泽瓦尔Sabzevar 	feud.Barony
	沙马特Shamat 	feud.Barony
}

func (c *内沙布尔NishapurCounty) BAkhlamad阿赫拉马德() feud.Barony {
	return c.阿赫拉马德Akhlamad
}
    
func (c *内沙布尔NishapurCounty) BChenaran恰纳兰() feud.Barony {
	return c.恰纳兰Chenaran
}
    
func (c *内沙布尔NishapurCounty) BDaregaz达尔加兹() feud.Barony {
	return c.达尔加兹Daregaz
}
    
func (c *内沙布尔NishapurCounty) BJajarm贾贾尔姆() feud.Barony {
	return c.贾贾尔姆Jajarm
}
    
func (c *内沙布尔NishapurCounty) BNishapur内沙布尔() feud.Barony {
	return c.内沙布尔Nishapur
}
    
func (c *内沙布尔NishapurCounty) BQuchan古昌() feud.Barony {
	return c.古昌Quchan
}
    
func (c *内沙布尔NishapurCounty) BSabzevar萨卜泽瓦尔() feud.Barony {
	return c.萨卜泽瓦尔Sabzevar
}
    
func (c *内沙布尔NishapurCounty) BShamat沙马特() feud.Barony {
	return c.沙马特Shamat
}
    
var CNishapur内沙布尔 NishapurCounty = &内沙布尔NishapurCounty{}

func init() {
	f := CNishapur内沙布尔.(*内沙布尔NishapurCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "634",
		Title:     "nishapur",
		TitleName: "内沙布尔",
		TitleCode: "c_nishapur",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿赫拉马德Akhlamad = BAkhlamad阿赫拉马德
	f.阿赫拉马德Akhlamad.SetParent(f)
	
	f.恰纳兰Chenaran = BChenaran恰纳兰
	f.恰纳兰Chenaran.SetParent(f)
	
	f.达尔加兹Daregaz = BDaregaz达尔加兹
	f.达尔加兹Daregaz.SetParent(f)
	
	f.贾贾尔姆Jajarm = BJajarm贾贾尔姆
	f.贾贾尔姆Jajarm.SetParent(f)
	
	f.内沙布尔Nishapur = BNishapur内沙布尔
	f.内沙布尔Nishapur.SetParent(f)
	
	f.古昌Quchan = BQuchan古昌
	f.古昌Quchan.SetParent(f)
	
	f.萨卜泽瓦尔Sabzevar = BSabzevar萨卜泽瓦尔
	f.萨卜泽瓦尔Sabzevar.SetParent(f)
	
	f.沙马特Shamat = BShamat沙马特
	f.沙马特Shamat.SetParent(f)
	
}
