package burkhan_buudai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Burkhan_buudaiCounty interface {
    feud.County
    BBiger比格尔() 	feud.Barony
    BBurkhan_buudai布尔汗布岱() 	feud.Barony
    BChandmani钱德曼() 	feud.Barony
    BErdene_burkhan额尔德尼() 	feud.Barony
    BKhaliun_burkhan哈里温() 	feud.Barony
    BTseel车勒() 	feud.Barony
    BTsogt朝格特() 	feud.Barony
}

type 布尔汗布岱Burkhan_buudaiCounty struct {
	feud.BaseCounty
	比格尔Biger 	feud.Barony
	布尔汗布岱Burkhan_buudai 	feud.Barony
	钱德曼Chandmani 	feud.Barony
	额尔德尼Erdene_burkhan 	feud.Barony
	哈里温Khaliun_burkhan 	feud.Barony
	车勒Tseel 	feud.Barony
	朝格特Tsogt 	feud.Barony
}

func (c *布尔汗布岱Burkhan_buudaiCounty) BBiger比格尔() feud.Barony {
	return c.比格尔Biger
}
    
func (c *布尔汗布岱Burkhan_buudaiCounty) BBurkhan_buudai布尔汗布岱() feud.Barony {
	return c.布尔汗布岱Burkhan_buudai
}
    
func (c *布尔汗布岱Burkhan_buudaiCounty) BChandmani钱德曼() feud.Barony {
	return c.钱德曼Chandmani
}
    
func (c *布尔汗布岱Burkhan_buudaiCounty) BErdene_burkhan额尔德尼() feud.Barony {
	return c.额尔德尼Erdene_burkhan
}
    
func (c *布尔汗布岱Burkhan_buudaiCounty) BKhaliun_burkhan哈里温() feud.Barony {
	return c.哈里温Khaliun_burkhan
}
    
func (c *布尔汗布岱Burkhan_buudaiCounty) BTseel车勒() feud.Barony {
	return c.车勒Tseel
}
    
func (c *布尔汗布岱Burkhan_buudaiCounty) BTsogt朝格特() feud.Barony {
	return c.朝格特Tsogt
}
    
var CBurkhan_buudai布尔汗布岱 Burkhan_buudaiCounty = &布尔汗布岱Burkhan_buudaiCounty{}

func init() {
	f := CBurkhan_buudai布尔汗布岱.(*布尔汗布岱Burkhan_buudaiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1916",
		Title:     "burkhan_buudai",
		TitleName: "布尔汗布岱",
		TitleCode: "c_burkhan_buudai",
		Baronies:  map[string]feud.Barony{},
	}

	f.比格尔Biger = BBiger比格尔
	f.比格尔Biger.SetParent(f)
	
	f.布尔汗布岱Burkhan_buudai = BBurkhan_buudai布尔汗布岱
	f.布尔汗布岱Burkhan_buudai.SetParent(f)
	
	f.钱德曼Chandmani = BChandmani钱德曼
	f.钱德曼Chandmani.SetParent(f)
	
	f.额尔德尼Erdene_burkhan = BErdene_burkhan额尔德尼
	f.额尔德尼Erdene_burkhan.SetParent(f)
	
	f.哈里温Khaliun_burkhan = BKhaliun_burkhan哈里温
	f.哈里温Khaliun_burkhan.SetParent(f)
	
	f.车勒Tseel = BTseel车勒
	f.车勒Tseel.SetParent(f)
	
	f.朝格特Tsogt = BTsogt朝格特
	f.朝格特Tsogt.SetParent(f)
	
}
