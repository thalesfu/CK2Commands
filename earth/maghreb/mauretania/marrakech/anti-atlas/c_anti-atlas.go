package anti-atlas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Anti-atlasCounty interface {
    feud.County
    BAgadir阿加迪尔() 	feud.Barony
    BAguedal阿盖达勒() 	feud.Barony
    BAitelaouni阿伊特阿乌里() 	feud.Barony
    BAkermoud阿克穆德() 	feud.Barony
    BArgana艾尔加奈() 	feud.Barony
    BEssaouira索维拉() 	feud.Barony
    BHaddraa哈德德拉() 	feud.Barony
    BIllrh安拉() 	feud.Barony
    BQumelaioun乌姆欧云() 	feud.Barony
    BSafi萨非() 	feud.Barony
}

type 萨非Anti-atlasCounty struct {
	feud.BaseCounty
	阿加迪尔Agadir 	feud.Barony
	阿盖达勒Aguedal 	feud.Barony
	阿伊特阿乌里Aitelaouni 	feud.Barony
	阿克穆德Akermoud 	feud.Barony
	艾尔加奈Argana 	feud.Barony
	索维拉Essaouira 	feud.Barony
	哈德德拉Haddraa 	feud.Barony
	安拉Illrh 	feud.Barony
	乌姆欧云Qumelaioun 	feud.Barony
	萨非Safi 	feud.Barony
}

func (c *萨非Anti-atlasCounty) BAgadir阿加迪尔() feud.Barony {
	return c.阿加迪尔Agadir
}
    
func (c *萨非Anti-atlasCounty) BAguedal阿盖达勒() feud.Barony {
	return c.阿盖达勒Aguedal
}
    
func (c *萨非Anti-atlasCounty) BAitelaouni阿伊特阿乌里() feud.Barony {
	return c.阿伊特阿乌里Aitelaouni
}
    
func (c *萨非Anti-atlasCounty) BAkermoud阿克穆德() feud.Barony {
	return c.阿克穆德Akermoud
}
    
func (c *萨非Anti-atlasCounty) BArgana艾尔加奈() feud.Barony {
	return c.艾尔加奈Argana
}
    
func (c *萨非Anti-atlasCounty) BEssaouira索维拉() feud.Barony {
	return c.索维拉Essaouira
}
    
func (c *萨非Anti-atlasCounty) BHaddraa哈德德拉() feud.Barony {
	return c.哈德德拉Haddraa
}
    
func (c *萨非Anti-atlasCounty) BIllrh安拉() feud.Barony {
	return c.安拉Illrh
}
    
func (c *萨非Anti-atlasCounty) BQumelaioun乌姆欧云() feud.Barony {
	return c.乌姆欧云Qumelaioun
}
    
func (c *萨非Anti-atlasCounty) BSafi萨非() feud.Barony {
	return c.萨非Safi
}
    
var CAnti-atlas萨非 Anti-atlasCounty = &萨非Anti-atlasCounty{}

func init() {
	f := CAnti-atlas萨非.(*萨非Anti-atlasCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "845",
		Title:     "anti-atlas",
		TitleName: "萨非",
		TitleCode: "c_anti-atlas",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿加迪尔Agadir = BAgadir阿加迪尔
	f.阿加迪尔Agadir.SetParent(f)
	
	f.阿盖达勒Aguedal = BAguedal阿盖达勒
	f.阿盖达勒Aguedal.SetParent(f)
	
	f.阿伊特阿乌里Aitelaouni = BAitelaouni阿伊特阿乌里
	f.阿伊特阿乌里Aitelaouni.SetParent(f)
	
	f.阿克穆德Akermoud = BAkermoud阿克穆德
	f.阿克穆德Akermoud.SetParent(f)
	
	f.艾尔加奈Argana = BArgana艾尔加奈
	f.艾尔加奈Argana.SetParent(f)
	
	f.索维拉Essaouira = BEssaouira索维拉
	f.索维拉Essaouira.SetParent(f)
	
	f.哈德德拉Haddraa = BHaddraa哈德德拉
	f.哈德德拉Haddraa.SetParent(f)
	
	f.安拉Illrh = BIllrh安拉
	f.安拉Illrh.SetParent(f)
	
	f.乌姆欧云Qumelaioun = BQumelaioun乌姆欧云
	f.乌姆欧云Qumelaioun.SetParent(f)
	
	f.萨非Safi = BSafi萨非
	f.萨非Safi.SetParent(f)
	
}
