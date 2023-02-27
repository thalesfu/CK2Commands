package votyaki

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VotyakiCounty interface {
    feud.County
    BAchit阿奇特() 	feud.Barony
    BArti阿尔季() 	feud.Barony
    BAtig阿季格() 	feud.Barony
    BBisert比谢尔季() 	feud.Barony
    BSarana萨拉纳() 	feud.Barony
    BShalya沙利亚() 	feud.Barony
    BShamary沙马雷() 	feud.Barony
    BUfimskiy乌菲姆斯基() 	feud.Barony
}

type 沃佳克VotyakiCounty struct {
	feud.BaseCounty
	阿奇特Achit 	feud.Barony
	阿尔季Arti 	feud.Barony
	阿季格Atig 	feud.Barony
	比谢尔季Bisert 	feud.Barony
	萨拉纳Sarana 	feud.Barony
	沙利亚Shalya 	feud.Barony
	沙马雷Shamary 	feud.Barony
	乌菲姆斯基Ufimskiy 	feud.Barony
}

func (c *沃佳克VotyakiCounty) BAchit阿奇特() feud.Barony {
	return c.阿奇特Achit
}
    
func (c *沃佳克VotyakiCounty) BArti阿尔季() feud.Barony {
	return c.阿尔季Arti
}
    
func (c *沃佳克VotyakiCounty) BAtig阿季格() feud.Barony {
	return c.阿季格Atig
}
    
func (c *沃佳克VotyakiCounty) BBisert比谢尔季() feud.Barony {
	return c.比谢尔季Bisert
}
    
func (c *沃佳克VotyakiCounty) BSarana萨拉纳() feud.Barony {
	return c.萨拉纳Sarana
}
    
func (c *沃佳克VotyakiCounty) BShalya沙利亚() feud.Barony {
	return c.沙利亚Shalya
}
    
func (c *沃佳克VotyakiCounty) BShamary沙马雷() feud.Barony {
	return c.沙马雷Shamary
}
    
func (c *沃佳克VotyakiCounty) BUfimskiy乌菲姆斯基() feud.Barony {
	return c.乌菲姆斯基Ufimskiy
}
    
var CVotyaki沃佳克 VotyakiCounty = &沃佳克VotyakiCounty{}

func init() {
	f := CVotyaki沃佳克.(*沃佳克VotyakiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "612",
		Title:     "votyaki",
		TitleName: "沃佳克",
		TitleCode: "c_votyaki",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿奇特Achit = BAchit阿奇特
	f.阿奇特Achit.SetParent(f)
	
	f.阿尔季Arti = BArti阿尔季
	f.阿尔季Arti.SetParent(f)
	
	f.阿季格Atig = BAtig阿季格
	f.阿季格Atig.SetParent(f)
	
	f.比谢尔季Bisert = BBisert比谢尔季
	f.比谢尔季Bisert.SetParent(f)
	
	f.萨拉纳Sarana = BSarana萨拉纳
	f.萨拉纳Sarana.SetParent(f)
	
	f.沙利亚Shalya = BShalya沙利亚
	f.沙利亚Shalya.SetParent(f)
	
	f.沙马雷Shamary = BShamary沙马雷
	f.沙马雷Shamary.SetParent(f)
	
	f.乌菲姆斯基Ufimskiy = BUfimskiy乌菲姆斯基
	f.乌菲姆斯基Ufimskiy.SetParent(f)
	
}
