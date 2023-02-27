package targu_jiu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Targu_jiuCounty interface {
    feud.County
    BBaia_de_arama巴亚德阿拉默() 	feud.Barony
    BBistrita比斯特里察() 	feud.Barony
    BPolovragi波洛夫拉吉() 	feud.Barony
    BTargu_jiu特尔古日乌() 	feud.Barony
    BTismana蒂斯马纳() 	feud.Barony
    BVisina维希纳() 	feud.Barony
    BVulcan武尔坎() 	feud.Barony
}

type 特尔古日乌Targu_jiuCounty struct {
	feud.BaseCounty
	巴亚德阿拉默Baia_de_arama 	feud.Barony
	比斯特里察Bistrita 	feud.Barony
	波洛夫拉吉Polovragi 	feud.Barony
	特尔古日乌Targu_jiu 	feud.Barony
	蒂斯马纳Tismana 	feud.Barony
	维希纳Visina 	feud.Barony
	武尔坎Vulcan 	feud.Barony
}

func (c *特尔古日乌Targu_jiuCounty) BBaia_de_arama巴亚德阿拉默() feud.Barony {
	return c.巴亚德阿拉默Baia_de_arama
}
    
func (c *特尔古日乌Targu_jiuCounty) BBistrita比斯特里察() feud.Barony {
	return c.比斯特里察Bistrita
}
    
func (c *特尔古日乌Targu_jiuCounty) BPolovragi波洛夫拉吉() feud.Barony {
	return c.波洛夫拉吉Polovragi
}
    
func (c *特尔古日乌Targu_jiuCounty) BTargu_jiu特尔古日乌() feud.Barony {
	return c.特尔古日乌Targu_jiu
}
    
func (c *特尔古日乌Targu_jiuCounty) BTismana蒂斯马纳() feud.Barony {
	return c.蒂斯马纳Tismana
}
    
func (c *特尔古日乌Targu_jiuCounty) BVisina维希纳() feud.Barony {
	return c.维希纳Visina
}
    
func (c *特尔古日乌Targu_jiuCounty) BVulcan武尔坎() feud.Barony {
	return c.武尔坎Vulcan
}
    
var CTargu_jiu特尔古日乌 Targu_jiuCounty = &特尔古日乌Targu_jiuCounty{}

func init() {
	f := CTargu_jiu特尔古日乌.(*特尔古日乌Targu_jiuCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1645",
		Title:     "targu_jiu",
		TitleName: "特尔古日乌",
		TitleCode: "c_targu_jiu",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴亚德阿拉默Baia_de_arama = BBaia_de_arama巴亚德阿拉默
	f.巴亚德阿拉默Baia_de_arama.SetParent(f)
	
	f.比斯特里察Bistrita = BBistrita比斯特里察
	f.比斯特里察Bistrita.SetParent(f)
	
	f.波洛夫拉吉Polovragi = BPolovragi波洛夫拉吉
	f.波洛夫拉吉Polovragi.SetParent(f)
	
	f.特尔古日乌Targu_jiu = BTargu_jiu特尔古日乌
	f.特尔古日乌Targu_jiu.SetParent(f)
	
	f.蒂斯马纳Tismana = BTismana蒂斯马纳
	f.蒂斯马纳Tismana.SetParent(f)
	
	f.维希纳Visina = BVisina维希纳
	f.维希纳Visina.SetParent(f)
	
	f.武尔坎Vulcan = BVulcan武尔坎
	f.武尔坎Vulcan.SetParent(f)
	
}
