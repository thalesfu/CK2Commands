package orsha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OrshaCounty interface {
    feud.County
    BBaran_orsha巴兰() 	feud.Barony
    BDubrowna杜布罗夫诺() 	feud.Barony
    BKrugale克鲁格韦() 	feud.Barony
    BOrsha奥尔沙() 	feud.Barony
    BShklow什克洛夫() 	feud.Barony
    BTalachyn塔拉钦() 	feud.Barony
    BZheleznyaki热列兹尼亚基() 	feud.Barony
}

type 奥尔沙OrshaCounty struct {
	feud.BaseCounty
	巴兰Baran_orsha 	feud.Barony
	杜布罗夫诺Dubrowna 	feud.Barony
	克鲁格韦Krugale 	feud.Barony
	奥尔沙Orsha 	feud.Barony
	什克洛夫Shklow 	feud.Barony
	塔拉钦Talachyn 	feud.Barony
	热列兹尼亚基Zheleznyaki 	feud.Barony
}

func (c *奥尔沙OrshaCounty) BBaran_orsha巴兰() feud.Barony {
	return c.巴兰Baran_orsha
}
    
func (c *奥尔沙OrshaCounty) BDubrowna杜布罗夫诺() feud.Barony {
	return c.杜布罗夫诺Dubrowna
}
    
func (c *奥尔沙OrshaCounty) BKrugale克鲁格韦() feud.Barony {
	return c.克鲁格韦Krugale
}
    
func (c *奥尔沙OrshaCounty) BOrsha奥尔沙() feud.Barony {
	return c.奥尔沙Orsha
}
    
func (c *奥尔沙OrshaCounty) BShklow什克洛夫() feud.Barony {
	return c.什克洛夫Shklow
}
    
func (c *奥尔沙OrshaCounty) BTalachyn塔拉钦() feud.Barony {
	return c.塔拉钦Talachyn
}
    
func (c *奥尔沙OrshaCounty) BZheleznyaki热列兹尼亚基() feud.Barony {
	return c.热列兹尼亚基Zheleznyaki
}
    
var COrsha奥尔沙 OrshaCounty = &奥尔沙OrshaCounty{}

func init() {
	f := COrsha奥尔沙.(*奥尔沙OrshaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "418",
		Title:     "orsha",
		TitleName: "奥尔沙",
		TitleCode: "c_orsha",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴兰Baran_orsha = BBaran_orsha巴兰
	f.巴兰Baran_orsha.SetParent(f)
	
	f.杜布罗夫诺Dubrowna = BDubrowna杜布罗夫诺
	f.杜布罗夫诺Dubrowna.SetParent(f)
	
	f.克鲁格韦Krugale = BKrugale克鲁格韦
	f.克鲁格韦Krugale.SetParent(f)
	
	f.奥尔沙Orsha = BOrsha奥尔沙
	f.奥尔沙Orsha.SetParent(f)
	
	f.什克洛夫Shklow = BShklow什克洛夫
	f.什克洛夫Shklow.SetParent(f)
	
	f.塔拉钦Talachyn = BTalachyn塔拉钦
	f.塔拉钦Talachyn.SetParent(f)
	
	f.热列兹尼亚基Zheleznyaki = BZheleznyaki热列兹尼亚基
	f.热列兹尼亚基Zheleznyaki.SetParent(f)
	
}
