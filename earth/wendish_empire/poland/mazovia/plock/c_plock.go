package plock

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PlockCounty interface {
    feud.County
    BCiechanow切哈努夫() 	feud.Barony
    BPlock普沃茨克() 	feud.Barony
    BPlonsk普翁斯克() 	feud.Barony
    BRozan鲁然() 	feud.Barony
    BSzrensk什伦斯克() 	feud.Barony
    BWyszogrod维绍格鲁德() 	feud.Barony
    BZacroczym扎克罗奇姆() 	feud.Barony
}

type 普沃茨克PlockCounty struct {
	feud.BaseCounty
	切哈努夫Ciechanow 	feud.Barony
	普沃茨克Plock 	feud.Barony
	普翁斯克Plonsk 	feud.Barony
	鲁然Rozan 	feud.Barony
	什伦斯克Szrensk 	feud.Barony
	维绍格鲁德Wyszogrod 	feud.Barony
	扎克罗奇姆Zacroczym 	feud.Barony
}

func (c *普沃茨克PlockCounty) BCiechanow切哈努夫() feud.Barony {
	return c.切哈努夫Ciechanow
}
    
func (c *普沃茨克PlockCounty) BPlock普沃茨克() feud.Barony {
	return c.普沃茨克Plock
}
    
func (c *普沃茨克PlockCounty) BPlonsk普翁斯克() feud.Barony {
	return c.普翁斯克Plonsk
}
    
func (c *普沃茨克PlockCounty) BRozan鲁然() feud.Barony {
	return c.鲁然Rozan
}
    
func (c *普沃茨克PlockCounty) BSzrensk什伦斯克() feud.Barony {
	return c.什伦斯克Szrensk
}
    
func (c *普沃茨克PlockCounty) BWyszogrod维绍格鲁德() feud.Barony {
	return c.维绍格鲁德Wyszogrod
}
    
func (c *普沃茨克PlockCounty) BZacroczym扎克罗奇姆() feud.Barony {
	return c.扎克罗奇姆Zacroczym
}
    
var CPlock普沃茨克 PlockCounty = &普沃茨克PlockCounty{}

func init() {
	f := CPlock普沃茨克.(*普沃茨克PlockCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "529",
		Title:     "plock",
		TitleName: "普沃茨克",
		TitleCode: "c_plock",
		Baronies:  map[string]feud.Barony{},
	}

	f.切哈努夫Ciechanow = BCiechanow切哈努夫
	f.切哈努夫Ciechanow.SetParent(f)
	
	f.普沃茨克Plock = BPlock普沃茨克
	f.普沃茨克Plock.SetParent(f)
	
	f.普翁斯克Plonsk = BPlonsk普翁斯克
	f.普翁斯克Plonsk.SetParent(f)
	
	f.鲁然Rozan = BRozan鲁然
	f.鲁然Rozan.SetParent(f)
	
	f.什伦斯克Szrensk = BSzrensk什伦斯克
	f.什伦斯克Szrensk.SetParent(f)
	
	f.维绍格鲁德Wyszogrod = BWyszogrod维绍格鲁德
	f.维绍格鲁德Wyszogrod.SetParent(f)
	
	f.扎克罗奇姆Zacroczym = BZacroczym扎克罗奇姆
	f.扎克罗奇姆Zacroczym.SetParent(f)
	
}
