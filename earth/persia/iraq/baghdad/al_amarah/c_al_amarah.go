package al_amarah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Al_amarahCounty interface {
    feud.County
    BAmarah阿马拉() 	feud.Barony
    BAzeeziaya阿泽吉阿亚() 	feud.Barony
    BBadra巴德拉() 	feud.Barony
    BHai哈里() 	feud.Barony
    BKutelamara库特阿玛拉() 	feud.Barony
    BSuwaira苏瓦拉() 	feud.Barony
    BWasit瓦斯特() 	feud.Barony
    BZurbatiyah祖尔拜亚() 	feud.Barony
}

type 阿马拉Al_amarahCounty struct {
	feud.BaseCounty
	阿马拉Amarah 	feud.Barony
	阿泽吉阿亚Azeeziaya 	feud.Barony
	巴德拉Badra 	feud.Barony
	哈里Hai 	feud.Barony
	库特阿玛拉Kutelamara 	feud.Barony
	苏瓦拉Suwaira 	feud.Barony
	瓦斯特Wasit 	feud.Barony
	祖尔拜亚Zurbatiyah 	feud.Barony
}

func (c *阿马拉Al_amarahCounty) BAmarah阿马拉() feud.Barony {
	return c.阿马拉Amarah
}
    
func (c *阿马拉Al_amarahCounty) BAzeeziaya阿泽吉阿亚() feud.Barony {
	return c.阿泽吉阿亚Azeeziaya
}
    
func (c *阿马拉Al_amarahCounty) BBadra巴德拉() feud.Barony {
	return c.巴德拉Badra
}
    
func (c *阿马拉Al_amarahCounty) BHai哈里() feud.Barony {
	return c.哈里Hai
}
    
func (c *阿马拉Al_amarahCounty) BKutelamara库特阿玛拉() feud.Barony {
	return c.库特阿玛拉Kutelamara
}
    
func (c *阿马拉Al_amarahCounty) BSuwaira苏瓦拉() feud.Barony {
	return c.苏瓦拉Suwaira
}
    
func (c *阿马拉Al_amarahCounty) BWasit瓦斯特() feud.Barony {
	return c.瓦斯特Wasit
}
    
func (c *阿马拉Al_amarahCounty) BZurbatiyah祖尔拜亚() feud.Barony {
	return c.祖尔拜亚Zurbatiyah
}
    
var CAl_amarah阿马拉 Al_amarahCounty = &阿马拉Al_amarahCounty{}

func init() {
	f := CAl_amarah阿马拉.(*阿马拉Al_amarahCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "690",
		Title:     "al_amarah",
		TitleName: "阿马拉",
		TitleCode: "c_al_amarah",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿马拉Amarah = BAmarah阿马拉
	f.阿马拉Amarah.SetParent(f)
	
	f.阿泽吉阿亚Azeeziaya = BAzeeziaya阿泽吉阿亚
	f.阿泽吉阿亚Azeeziaya.SetParent(f)
	
	f.巴德拉Badra = BBadra巴德拉
	f.巴德拉Badra.SetParent(f)
	
	f.哈里Hai = BHai哈里
	f.哈里Hai.SetParent(f)
	
	f.库特阿玛拉Kutelamara = BKutelamara库特阿玛拉
	f.库特阿玛拉Kutelamara.SetParent(f)
	
	f.苏瓦拉Suwaira = BSuwaira苏瓦拉
	f.苏瓦拉Suwaira.SetParent(f)
	
	f.瓦斯特Wasit = BWasit瓦斯特
	f.瓦斯特Wasit.SetParent(f)
	
	f.祖尔拜亚Zurbatiyah = BZurbatiyah祖尔拜亚
	f.祖尔拜亚Zurbatiyah.SetParent(f)
	
}
