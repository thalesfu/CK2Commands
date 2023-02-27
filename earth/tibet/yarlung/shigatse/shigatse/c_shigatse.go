package shigatse

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ShigatseCounty interface {
    feud.County
    BGongon_lhakhang贵恩拉康() 	feud.Barony
    BNamling南木林() 	feud.Barony
    BSamzhubze三竹节() 	feud.Barony
    BShalu夏鲁() 	feud.Barony
    BShigatse日喀则() 	feud.Barony
    BTashilhunpo扎什伦布() 	feud.Barony
    BYungdrungling雍仲林() 	feud.Barony
}

type 日喀则ShigatseCounty struct {
	feud.BaseCounty
	贵恩拉康Gongon_lhakhang 	feud.Barony
	南木林Namling 	feud.Barony
	三竹节Samzhubze 	feud.Barony
	夏鲁Shalu 	feud.Barony
	日喀则Shigatse 	feud.Barony
	扎什伦布Tashilhunpo 	feud.Barony
	雍仲林Yungdrungling 	feud.Barony
}

func (c *日喀则ShigatseCounty) BGongon_lhakhang贵恩拉康() feud.Barony {
	return c.贵恩拉康Gongon_lhakhang
}
    
func (c *日喀则ShigatseCounty) BNamling南木林() feud.Barony {
	return c.南木林Namling
}
    
func (c *日喀则ShigatseCounty) BSamzhubze三竹节() feud.Barony {
	return c.三竹节Samzhubze
}
    
func (c *日喀则ShigatseCounty) BShalu夏鲁() feud.Barony {
	return c.夏鲁Shalu
}
    
func (c *日喀则ShigatseCounty) BShigatse日喀则() feud.Barony {
	return c.日喀则Shigatse
}
    
func (c *日喀则ShigatseCounty) BTashilhunpo扎什伦布() feud.Barony {
	return c.扎什伦布Tashilhunpo
}
    
func (c *日喀则ShigatseCounty) BYungdrungling雍仲林() feud.Barony {
	return c.雍仲林Yungdrungling
}
    
var CShigatse日喀则 ShigatseCounty = &日喀则ShigatseCounty{}

func init() {
	f := CShigatse日喀则.(*日喀则ShigatseCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1499",
		Title:     "shigatse",
		TitleName: "日喀则",
		TitleCode: "c_shigatse",
		Baronies:  map[string]feud.Barony{},
	}

	f.贵恩拉康Gongon_lhakhang = BGongon_lhakhang贵恩拉康
	f.贵恩拉康Gongon_lhakhang.SetParent(f)
	
	f.南木林Namling = BNamling南木林
	f.南木林Namling.SetParent(f)
	
	f.三竹节Samzhubze = BSamzhubze三竹节
	f.三竹节Samzhubze.SetParent(f)
	
	f.夏鲁Shalu = BShalu夏鲁
	f.夏鲁Shalu.SetParent(f)
	
	f.日喀则Shigatse = BShigatse日喀则
	f.日喀则Shigatse.SetParent(f)
	
	f.扎什伦布Tashilhunpo = BTashilhunpo扎什伦布
	f.扎什伦布Tashilhunpo.SetParent(f)
	
	f.雍仲林Yungdrungling = BYungdrungling雍仲林
	f.雍仲林Yungdrungling.SetParent(f)
	
}
