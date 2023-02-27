package cieszyn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CieszynCounty interface {
    feud.County
    BBielsko别尔斯科() 	feud.Barony
    BCieszyn切申() 	feud.Barony
    BOswiecim奥斯威辛() 	feud.Barony
    BPszczyna普什奇纳() 	feud.Barony
    BSkoczow斯科丘夫() 	feud.Barony
    BZator扎托尔() 	feud.Barony
    BZywiec日维茨() 	feud.Barony
}

type 切申CieszynCounty struct {
	feud.BaseCounty
	别尔斯科Bielsko 	feud.Barony
	切申Cieszyn 	feud.Barony
	奥斯威辛Oswiecim 	feud.Barony
	普什奇纳Pszczyna 	feud.Barony
	斯科丘夫Skoczow 	feud.Barony
	扎托尔Zator 	feud.Barony
	日维茨Zywiec 	feud.Barony
}

func (c *切申CieszynCounty) BBielsko别尔斯科() feud.Barony {
	return c.别尔斯科Bielsko
}
    
func (c *切申CieszynCounty) BCieszyn切申() feud.Barony {
	return c.切申Cieszyn
}
    
func (c *切申CieszynCounty) BOswiecim奥斯威辛() feud.Barony {
	return c.奥斯威辛Oswiecim
}
    
func (c *切申CieszynCounty) BPszczyna普什奇纳() feud.Barony {
	return c.普什奇纳Pszczyna
}
    
func (c *切申CieszynCounty) BSkoczow斯科丘夫() feud.Barony {
	return c.斯科丘夫Skoczow
}
    
func (c *切申CieszynCounty) BZator扎托尔() feud.Barony {
	return c.扎托尔Zator
}
    
func (c *切申CieszynCounty) BZywiec日维茨() feud.Barony {
	return c.日维茨Zywiec
}
    
var CCieszyn切申 CieszynCounty = &切申CieszynCounty{}

func init() {
	f := CCieszyn切申.(*切申CieszynCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "526",
		Title:     "cieszyn",
		TitleName: "切申",
		TitleCode: "c_cieszyn",
		Baronies:  map[string]feud.Barony{},
	}

	f.别尔斯科Bielsko = BBielsko别尔斯科
	f.别尔斯科Bielsko.SetParent(f)
	
	f.切申Cieszyn = BCieszyn切申
	f.切申Cieszyn.SetParent(f)
	
	f.奥斯威辛Oswiecim = BOswiecim奥斯威辛
	f.奥斯威辛Oswiecim.SetParent(f)
	
	f.普什奇纳Pszczyna = BPszczyna普什奇纳
	f.普什奇纳Pszczyna.SetParent(f)
	
	f.斯科丘夫Skoczow = BSkoczow斯科丘夫
	f.斯科丘夫Skoczow.SetParent(f)
	
	f.扎托尔Zator = BZator扎托尔
	f.扎托尔Zator.SetParent(f)
	
	f.日维茨Zywiec = BZywiec日维茨
	f.日维茨Zywiec.SetParent(f)
	
}
