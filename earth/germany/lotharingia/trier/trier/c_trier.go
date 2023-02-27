package trier

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TrierCounty interface {
    feud.County
    BAndernach安德纳赫() 	feud.Barony
    BBitburg比特堡() 	feud.Barony
    BCoblenz科布伦茨() 	feud.Barony
    BGerolstein盖罗尔施泰因() 	feud.Barony
    BLaach拉赫() 	feud.Barony
    BSponheim施蓬海姆() 	feud.Barony
    BTrier特里尔() 	feud.Barony
    BWittlich维特利希() 	feud.Barony
}

type 特里尔TrierCounty struct {
	feud.BaseCounty
	安德纳赫Andernach 	feud.Barony
	比特堡Bitburg 	feud.Barony
	科布伦茨Coblenz 	feud.Barony
	盖罗尔施泰因Gerolstein 	feud.Barony
	拉赫Laach 	feud.Barony
	施蓬海姆Sponheim 	feud.Barony
	特里尔Trier 	feud.Barony
	维特利希Wittlich 	feud.Barony
}

func (c *特里尔TrierCounty) BAndernach安德纳赫() feud.Barony {
	return c.安德纳赫Andernach
}
    
func (c *特里尔TrierCounty) BBitburg比特堡() feud.Barony {
	return c.比特堡Bitburg
}
    
func (c *特里尔TrierCounty) BCoblenz科布伦茨() feud.Barony {
	return c.科布伦茨Coblenz
}
    
func (c *特里尔TrierCounty) BGerolstein盖罗尔施泰因() feud.Barony {
	return c.盖罗尔施泰因Gerolstein
}
    
func (c *特里尔TrierCounty) BLaach拉赫() feud.Barony {
	return c.拉赫Laach
}
    
func (c *特里尔TrierCounty) BSponheim施蓬海姆() feud.Barony {
	return c.施蓬海姆Sponheim
}
    
func (c *特里尔TrierCounty) BTrier特里尔() feud.Barony {
	return c.特里尔Trier
}
    
func (c *特里尔TrierCounty) BWittlich维特利希() feud.Barony {
	return c.维特利希Wittlich
}
    
var CTrier特里尔 TrierCounty = &特里尔TrierCounty{}

func init() {
	f := CTrier特里尔.(*特里尔TrierCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "118",
		Title:     "trier",
		TitleName: "特里尔",
		TitleCode: "c_trier",
		Baronies:  map[string]feud.Barony{},
	}

	f.安德纳赫Andernach = BAndernach安德纳赫
	f.安德纳赫Andernach.SetParent(f)
	
	f.比特堡Bitburg = BBitburg比特堡
	f.比特堡Bitburg.SetParent(f)
	
	f.科布伦茨Coblenz = BCoblenz科布伦茨
	f.科布伦茨Coblenz.SetParent(f)
	
	f.盖罗尔施泰因Gerolstein = BGerolstein盖罗尔施泰因
	f.盖罗尔施泰因Gerolstein.SetParent(f)
	
	f.拉赫Laach = BLaach拉赫
	f.拉赫Laach.SetParent(f)
	
	f.施蓬海姆Sponheim = BSponheim施蓬海姆
	f.施蓬海姆Sponheim.SetParent(f)
	
	f.特里尔Trier = BTrier特里尔
	f.特里尔Trier.SetParent(f)
	
	f.维特利希Wittlich = BWittlich维特利希
	f.维特利希Wittlich.SetParent(f)
	
}
