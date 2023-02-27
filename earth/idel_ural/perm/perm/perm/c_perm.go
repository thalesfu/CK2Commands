package perm

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PermCounty interface {
    feud.County
    BBiser比谢尔() 	feud.Barony
    BChemuska切穆斯卡() 	feud.Barony
    BCherdyn切尔登() 	feud.Barony
    BGorodki格罗德基() 	feud.Barony
    BLysva雷西瓦() 	feud.Barony
    BPerm彼尔姆() 	feud.Barony
    BPerym佩雷姆() 	feud.Barony
    BYagoshikha亚戈希哈() 	feud.Barony
}

type 彼尔姆PermCounty struct {
	feud.BaseCounty
	比谢尔Biser 	feud.Barony
	切穆斯卡Chemuska 	feud.Barony
	切尔登Cherdyn 	feud.Barony
	格罗德基Gorodki 	feud.Barony
	雷西瓦Lysva 	feud.Barony
	彼尔姆Perm 	feud.Barony
	佩雷姆Perym 	feud.Barony
	亚戈希哈Yagoshikha 	feud.Barony
}

func (c *彼尔姆PermCounty) BBiser比谢尔() feud.Barony {
	return c.比谢尔Biser
}
    
func (c *彼尔姆PermCounty) BChemuska切穆斯卡() feud.Barony {
	return c.切穆斯卡Chemuska
}
    
func (c *彼尔姆PermCounty) BCherdyn切尔登() feud.Barony {
	return c.切尔登Cherdyn
}
    
func (c *彼尔姆PermCounty) BGorodki格罗德基() feud.Barony {
	return c.格罗德基Gorodki
}
    
func (c *彼尔姆PermCounty) BLysva雷西瓦() feud.Barony {
	return c.雷西瓦Lysva
}
    
func (c *彼尔姆PermCounty) BPerm彼尔姆() feud.Barony {
	return c.彼尔姆Perm
}
    
func (c *彼尔姆PermCounty) BPerym佩雷姆() feud.Barony {
	return c.佩雷姆Perym
}
    
func (c *彼尔姆PermCounty) BYagoshikha亚戈希哈() feud.Barony {
	return c.亚戈希哈Yagoshikha
}
    
var CPerm彼尔姆 PermCounty = &彼尔姆PermCounty{}

func init() {
	f := CPerm彼尔姆.(*彼尔姆PermCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "886",
		Title:     "perm",
		TitleName: "彼尔姆",
		TitleCode: "c_perm",
		Baronies:  map[string]feud.Barony{},
	}

	f.比谢尔Biser = BBiser比谢尔
	f.比谢尔Biser.SetParent(f)
	
	f.切穆斯卡Chemuska = BChemuska切穆斯卡
	f.切穆斯卡Chemuska.SetParent(f)
	
	f.切尔登Cherdyn = BCherdyn切尔登
	f.切尔登Cherdyn.SetParent(f)
	
	f.格罗德基Gorodki = BGorodki格罗德基
	f.格罗德基Gorodki.SetParent(f)
	
	f.雷西瓦Lysva = BLysva雷西瓦
	f.雷西瓦Lysva.SetParent(f)
	
	f.彼尔姆Perm = BPerm彼尔姆
	f.彼尔姆Perm.SetParent(f)
	
	f.佩雷姆Perym = BPerym佩雷姆
	f.佩雷姆Perym.SetParent(f)
	
	f.亚戈希哈Yagoshikha = BYagoshikha亚戈希哈
	f.亚戈希哈Yagoshikha.SetParent(f)
	
}
