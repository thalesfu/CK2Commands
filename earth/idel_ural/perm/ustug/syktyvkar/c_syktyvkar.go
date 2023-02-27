package syktyvkar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SyktyvkarCounty interface {
    feud.County
    BChoy乔伊() 	feud.Barony
    BKoyty科伊特() 	feud.Barony
    BPazhga帕日加() 	feud.Barony
    BSyktyvkar瑟克特夫卡尔() 	feud.Barony
    BSysola瑟索拉() 	feud.Barony
    BVyl_gort维利戈尔特() 	feud.Barony
    BZelenets泽列涅茨() 	feud.Barony
}

type 瑟克特夫卡尔SyktyvkarCounty struct {
	feud.BaseCounty
	乔伊Choy 	feud.Barony
	科伊特Koyty 	feud.Barony
	帕日加Pazhga 	feud.Barony
	瑟克特夫卡尔Syktyvkar 	feud.Barony
	瑟索拉Sysola 	feud.Barony
	维利戈尔特Vyl_gort 	feud.Barony
	泽列涅茨Zelenets 	feud.Barony
}

func (c *瑟克特夫卡尔SyktyvkarCounty) BChoy乔伊() feud.Barony {
	return c.乔伊Choy
}
    
func (c *瑟克特夫卡尔SyktyvkarCounty) BKoyty科伊特() feud.Barony {
	return c.科伊特Koyty
}
    
func (c *瑟克特夫卡尔SyktyvkarCounty) BPazhga帕日加() feud.Barony {
	return c.帕日加Pazhga
}
    
func (c *瑟克特夫卡尔SyktyvkarCounty) BSyktyvkar瑟克特夫卡尔() feud.Barony {
	return c.瑟克特夫卡尔Syktyvkar
}
    
func (c *瑟克特夫卡尔SyktyvkarCounty) BSysola瑟索拉() feud.Barony {
	return c.瑟索拉Sysola
}
    
func (c *瑟克特夫卡尔SyktyvkarCounty) BVyl_gort维利戈尔特() feud.Barony {
	return c.维利戈尔特Vyl_gort
}
    
func (c *瑟克特夫卡尔SyktyvkarCounty) BZelenets泽列涅茨() feud.Barony {
	return c.泽列涅茨Zelenets
}
    
var CSyktyvkar瑟克特夫卡尔 SyktyvkarCounty = &瑟克特夫卡尔SyktyvkarCounty{}

func init() {
	f := CSyktyvkar瑟克特夫卡尔.(*瑟克特夫卡尔SyktyvkarCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1817",
		Title:     "syktyvkar",
		TitleName: "瑟克特夫卡尔",
		TitleCode: "c_syktyvkar",
		Baronies:  map[string]feud.Barony{},
	}

	f.乔伊Choy = BChoy乔伊
	f.乔伊Choy.SetParent(f)
	
	f.科伊特Koyty = BKoyty科伊特
	f.科伊特Koyty.SetParent(f)
	
	f.帕日加Pazhga = BPazhga帕日加
	f.帕日加Pazhga.SetParent(f)
	
	f.瑟克特夫卡尔Syktyvkar = BSyktyvkar瑟克特夫卡尔
	f.瑟克特夫卡尔Syktyvkar.SetParent(f)
	
	f.瑟索拉Sysola = BSysola瑟索拉
	f.瑟索拉Sysola.SetParent(f)
	
	f.维利戈尔特Vyl_gort = BVyl_gort维利戈尔特
	f.维利戈尔特Vyl_gort.SetParent(f)
	
	f.泽列涅茨Zelenets = BZelenets泽列涅茨
	f.泽列涅茨Zelenets.SetParent(f)
	
}
