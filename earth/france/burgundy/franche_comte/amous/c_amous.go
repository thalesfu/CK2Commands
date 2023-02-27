package amous

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AmousCounty interface {
    feud.County
    BAmaous阿玛乌() 	feud.Barony
    BChamaves萨马维斯() 	feud.Barony
    BDole多勒() 	feud.Barony
    BGray格雷() 	feud.Barony
    BJouhe茹埃() 	feud.Barony
    BNotredame诺特尔达姆() 	feud.Barony
    BSaintylie圣伊利() 	feud.Barony
}

type 阿穆AmousCounty struct {
	feud.BaseCounty
	阿玛乌Amaous 	feud.Barony
	萨马维斯Chamaves 	feud.Barony
	多勒Dole 	feud.Barony
	格雷Gray 	feud.Barony
	茹埃Jouhe 	feud.Barony
	诺特尔达姆Notredame 	feud.Barony
	圣伊利Saintylie 	feud.Barony
}

func (c *阿穆AmousCounty) BAmaous阿玛乌() feud.Barony {
	return c.阿玛乌Amaous
}
    
func (c *阿穆AmousCounty) BChamaves萨马维斯() feud.Barony {
	return c.萨马维斯Chamaves
}
    
func (c *阿穆AmousCounty) BDole多勒() feud.Barony {
	return c.多勒Dole
}
    
func (c *阿穆AmousCounty) BGray格雷() feud.Barony {
	return c.格雷Gray
}
    
func (c *阿穆AmousCounty) BJouhe茹埃() feud.Barony {
	return c.茹埃Jouhe
}
    
func (c *阿穆AmousCounty) BNotredame诺特尔达姆() feud.Barony {
	return c.诺特尔达姆Notredame
}
    
func (c *阿穆AmousCounty) BSaintylie圣伊利() feud.Barony {
	return c.圣伊利Saintylie
}
    
var CAmous阿穆 AmousCounty = &阿穆AmousCounty{}

func init() {
	f := CAmous阿穆.(*阿穆AmousCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1615",
		Title:     "amous",
		TitleName: "阿穆",
		TitleCode: "c_amous",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿玛乌Amaous = BAmaous阿玛乌
	f.阿玛乌Amaous.SetParent(f)
	
	f.萨马维斯Chamaves = BChamaves萨马维斯
	f.萨马维斯Chamaves.SetParent(f)
	
	f.多勒Dole = BDole多勒
	f.多勒Dole.SetParent(f)
	
	f.格雷Gray = BGray格雷
	f.格雷Gray.SetParent(f)
	
	f.茹埃Jouhe = BJouhe茹埃
	f.茹埃Jouhe.SetParent(f)
	
	f.诺特尔达姆Notredame = BNotredame诺特尔达姆
	f.诺特尔达姆Notredame.SetParent(f)
	
	f.圣伊利Saintylie = BSaintylie圣伊利
	f.圣伊利Saintylie.SetParent(f)
	
}
