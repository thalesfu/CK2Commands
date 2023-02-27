package tihama

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TihamaCounty interface {
    feud.County
    BAbs阿布斯() 	feud.Barony
    BAs_suqaiq苏盖格() 	feud.Barony
    BAttar阿塔尔() 	feud.Barony
    BBays拜斯() 	feud.Barony
    BHali哈利() 	feud.Barony
    BHamr哈穆伦() 	feud.Barony
    BHarad哈拉德() 	feud.Barony
}

type 帖哈麦TihamaCounty struct {
	feud.BaseCounty
	阿布斯Abs 	feud.Barony
	苏盖格As_suqaiq 	feud.Barony
	阿塔尔Attar 	feud.Barony
	拜斯Bays 	feud.Barony
	哈利Hali 	feud.Barony
	哈穆伦Hamr 	feud.Barony
	哈拉德Harad 	feud.Barony
}

func (c *帖哈麦TihamaCounty) BAbs阿布斯() feud.Barony {
	return c.阿布斯Abs
}
    
func (c *帖哈麦TihamaCounty) BAs_suqaiq苏盖格() feud.Barony {
	return c.苏盖格As_suqaiq
}
    
func (c *帖哈麦TihamaCounty) BAttar阿塔尔() feud.Barony {
	return c.阿塔尔Attar
}
    
func (c *帖哈麦TihamaCounty) BBays拜斯() feud.Barony {
	return c.拜斯Bays
}
    
func (c *帖哈麦TihamaCounty) BHali哈利() feud.Barony {
	return c.哈利Hali
}
    
func (c *帖哈麦TihamaCounty) BHamr哈穆伦() feud.Barony {
	return c.哈穆伦Hamr
}
    
func (c *帖哈麦TihamaCounty) BHarad哈拉德() feud.Barony {
	return c.哈拉德Harad
}
    
var CTihama帖哈麦 TihamaCounty = &帖哈麦TihamaCounty{}

func init() {
	f := CTihama帖哈麦.(*帖哈麦TihamaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1535",
		Title:     "tihama",
		TitleName: "帖哈麦",
		TitleCode: "c_tihama",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿布斯Abs = BAbs阿布斯
	f.阿布斯Abs.SetParent(f)
	
	f.苏盖格As_suqaiq = BAs_suqaiq苏盖格
	f.苏盖格As_suqaiq.SetParent(f)
	
	f.阿塔尔Attar = BAttar阿塔尔
	f.阿塔尔Attar.SetParent(f)
	
	f.拜斯Bays = BBays拜斯
	f.拜斯Bays.SetParent(f)
	
	f.哈利Hali = BHali哈利
	f.哈利Hali.SetParent(f)
	
	f.哈穆伦Hamr = BHamr哈穆伦
	f.哈穆伦Hamr.SetParent(f)
	
	f.哈拉德Harad = BHarad哈拉德
	f.哈拉德Harad.SetParent(f)
	
}
