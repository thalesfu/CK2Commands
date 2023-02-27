package mstislavl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MstislavlCounty interface {
    feud.County
    BChachersk切切尔斯克() 	feud.Barony
    BDrybin德里宾() 	feud.Barony
    BHorki戈尔基() 	feud.Barony
    BKrychaw克里切夫() 	feud.Barony
    BMstitslavl姆斯季斯拉夫尔() 	feud.Barony
    BPropoysk普罗波伊斯克() 	feud.Barony
    BRahacou罗加乔夫() 	feud.Barony
}

type 姆斯季斯拉夫尔MstislavlCounty struct {
	feud.BaseCounty
	切切尔斯克Chachersk 	feud.Barony
	德里宾Drybin 	feud.Barony
	戈尔基Horki 	feud.Barony
	克里切夫Krychaw 	feud.Barony
	姆斯季斯拉夫尔Mstitslavl 	feud.Barony
	普罗波伊斯克Propoysk 	feud.Barony
	罗加乔夫Rahacou 	feud.Barony
}

func (c *姆斯季斯拉夫尔MstislavlCounty) BChachersk切切尔斯克() feud.Barony {
	return c.切切尔斯克Chachersk
}
    
func (c *姆斯季斯拉夫尔MstislavlCounty) BDrybin德里宾() feud.Barony {
	return c.德里宾Drybin
}
    
func (c *姆斯季斯拉夫尔MstislavlCounty) BHorki戈尔基() feud.Barony {
	return c.戈尔基Horki
}
    
func (c *姆斯季斯拉夫尔MstislavlCounty) BKrychaw克里切夫() feud.Barony {
	return c.克里切夫Krychaw
}
    
func (c *姆斯季斯拉夫尔MstislavlCounty) BMstitslavl姆斯季斯拉夫尔() feud.Barony {
	return c.姆斯季斯拉夫尔Mstitslavl
}
    
func (c *姆斯季斯拉夫尔MstislavlCounty) BPropoysk普罗波伊斯克() feud.Barony {
	return c.普罗波伊斯克Propoysk
}
    
func (c *姆斯季斯拉夫尔MstislavlCounty) BRahacou罗加乔夫() feud.Barony {
	return c.罗加乔夫Rahacou
}
    
var CMstislavl姆斯季斯拉夫尔 MstislavlCounty = &姆斯季斯拉夫尔MstislavlCounty{}

func init() {
	f := CMstislavl姆斯季斯拉夫尔.(*姆斯季斯拉夫尔MstislavlCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "551",
		Title:     "mstislavl",
		TitleName: "姆斯季斯拉夫尔",
		TitleCode: "c_mstislavl",
		Baronies:  map[string]feud.Barony{},
	}

	f.切切尔斯克Chachersk = BChachersk切切尔斯克
	f.切切尔斯克Chachersk.SetParent(f)
	
	f.德里宾Drybin = BDrybin德里宾
	f.德里宾Drybin.SetParent(f)
	
	f.戈尔基Horki = BHorki戈尔基
	f.戈尔基Horki.SetParent(f)
	
	f.克里切夫Krychaw = BKrychaw克里切夫
	f.克里切夫Krychaw.SetParent(f)
	
	f.姆斯季斯拉夫尔Mstitslavl = BMstitslavl姆斯季斯拉夫尔
	f.姆斯季斯拉夫尔Mstitslavl.SetParent(f)
	
	f.普罗波伊斯克Propoysk = BPropoysk普罗波伊斯克
	f.普罗波伊斯克Propoysk.SetParent(f)
	
	f.罗加乔夫Rahacou = BRahacou罗加乔夫
	f.罗加乔夫Rahacou.SetParent(f)
	
}
