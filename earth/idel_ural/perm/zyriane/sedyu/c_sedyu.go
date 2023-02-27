package sedyu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SedyuCounty interface {
    feud.County
    BAkim阿基姆() 	feud.Barony
    BAyuva_sedyu阿尤瓦() 	feud.Barony
    BKedva克德瓦() 	feud.Barony
    BSebys谢贝西() 	feud.Barony
    BSedyu谢季尤() 	feud.Barony
    BUkhta_sedyu乌赫塔() 	feud.Barony
    BYuger尤格尔() 	feud.Barony
}

type 谢季尤SedyuCounty struct {
	feud.BaseCounty
	阿基姆Akim 	feud.Barony
	阿尤瓦Ayuva_sedyu 	feud.Barony
	克德瓦Kedva 	feud.Barony
	谢贝西Sebys 	feud.Barony
	谢季尤Sedyu 	feud.Barony
	乌赫塔Ukhta_sedyu 	feud.Barony
	尤格尔Yuger 	feud.Barony
}

func (c *谢季尤SedyuCounty) BAkim阿基姆() feud.Barony {
	return c.阿基姆Akim
}
    
func (c *谢季尤SedyuCounty) BAyuva_sedyu阿尤瓦() feud.Barony {
	return c.阿尤瓦Ayuva_sedyu
}
    
func (c *谢季尤SedyuCounty) BKedva克德瓦() feud.Barony {
	return c.克德瓦Kedva
}
    
func (c *谢季尤SedyuCounty) BSebys谢贝西() feud.Barony {
	return c.谢贝西Sebys
}
    
func (c *谢季尤SedyuCounty) BSedyu谢季尤() feud.Barony {
	return c.谢季尤Sedyu
}
    
func (c *谢季尤SedyuCounty) BUkhta_sedyu乌赫塔() feud.Barony {
	return c.乌赫塔Ukhta_sedyu
}
    
func (c *谢季尤SedyuCounty) BYuger尤格尔() feud.Barony {
	return c.尤格尔Yuger
}
    
var CSedyu谢季尤 SedyuCounty = &谢季尤SedyuCounty{}

func init() {
	f := CSedyu谢季尤.(*谢季尤SedyuCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1840",
		Title:     "sedyu",
		TitleName: "谢季尤",
		TitleCode: "c_sedyu",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿基姆Akim = BAkim阿基姆
	f.阿基姆Akim.SetParent(f)
	
	f.阿尤瓦Ayuva_sedyu = BAyuva_sedyu阿尤瓦
	f.阿尤瓦Ayuva_sedyu.SetParent(f)
	
	f.克德瓦Kedva = BKedva克德瓦
	f.克德瓦Kedva.SetParent(f)
	
	f.谢贝西Sebys = BSebys谢贝西
	f.谢贝西Sebys.SetParent(f)
	
	f.谢季尤Sedyu = BSedyu谢季尤
	f.谢季尤Sedyu.SetParent(f)
	
	f.乌赫塔Ukhta_sedyu = BUkhta_sedyu乌赫塔
	f.乌赫塔Ukhta_sedyu.SetParent(f)
	
	f.尤格尔Yuger = BYuger尤格尔
	f.尤格尔Yuger.SetParent(f)
	
}
