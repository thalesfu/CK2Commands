package bath

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BathCounty interface {
    feud.County
    BBath巴斯() 	feud.Barony
    BBristol布里斯托() 	feud.Barony
    BBurnham伯纳姆() 	feud.Barony
    BCadbury卡德伯里() 	feud.Barony
    BCastle_batch巴奇() 	feud.Barony
    BCleeve克利夫() 	feud.Barony
    BWells韦尔斯() 	feud.Barony
}

type 巴斯BathCounty struct {
	feud.BaseCounty
	巴斯Bath 	feud.Barony
	布里斯托Bristol 	feud.Barony
	伯纳姆Burnham 	feud.Barony
	卡德伯里Cadbury 	feud.Barony
	巴奇Castle_batch 	feud.Barony
	克利夫Cleeve 	feud.Barony
	韦尔斯Wells 	feud.Barony
}

func (c *巴斯BathCounty) BBath巴斯() feud.Barony {
	return c.巴斯Bath
}
    
func (c *巴斯BathCounty) BBristol布里斯托() feud.Barony {
	return c.布里斯托Bristol
}
    
func (c *巴斯BathCounty) BBurnham伯纳姆() feud.Barony {
	return c.伯纳姆Burnham
}
    
func (c *巴斯BathCounty) BCadbury卡德伯里() feud.Barony {
	return c.卡德伯里Cadbury
}
    
func (c *巴斯BathCounty) BCastle_batch巴奇() feud.Barony {
	return c.巴奇Castle_batch
}
    
func (c *巴斯BathCounty) BCleeve克利夫() feud.Barony {
	return c.克利夫Cleeve
}
    
func (c *巴斯BathCounty) BWells韦尔斯() feud.Barony {
	return c.韦尔斯Wells
}
    
var CBath巴斯 BathCounty = &巴斯BathCounty{}

func init() {
	f := CBath巴斯.(*巴斯BathCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1944",
		Title:     "bath",
		TitleName: "巴斯",
		TitleCode: "c_bath",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴斯Bath = BBath巴斯
	f.巴斯Bath.SetParent(f)
	
	f.布里斯托Bristol = BBristol布里斯托
	f.布里斯托Bristol.SetParent(f)
	
	f.伯纳姆Burnham = BBurnham伯纳姆
	f.伯纳姆Burnham.SetParent(f)
	
	f.卡德伯里Cadbury = BCadbury卡德伯里
	f.卡德伯里Cadbury.SetParent(f)
	
	f.巴奇Castle_batch = BCastle_batch巴奇
	f.巴奇Castle_batch.SetParent(f)
	
	f.克利夫Cleeve = BCleeve克利夫
	f.克利夫Cleeve.SetParent(f)
	
	f.韦尔斯Wells = BWells韦尔斯
	f.韦尔斯Wells.SetParent(f)
	
}
