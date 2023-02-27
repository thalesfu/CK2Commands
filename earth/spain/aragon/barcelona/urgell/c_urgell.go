package urgell

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type UrgellCounty interface {
    feud.County
    BElpuidesegur厄尔贝塞居尔() 	feud.Barony
    BPallars帕拉尔斯() 	feud.Barony
    BPuigcerda普奇塞达() 	feud.Barony
    BSuert苏埃尔特() 	feud.Barony
    BTremp特伦普() 	feud.Barony
    BUrgell乌尔赫尔() 	feud.Barony
    BValledebohi巴利德沃伊() 	feud.Barony
    BViella别利亚() 	feud.Barony
}

type 乌尔赫尔UrgellCounty struct {
	feud.BaseCounty
	厄尔贝塞居尔Elpuidesegur 	feud.Barony
	帕拉尔斯Pallars 	feud.Barony
	普奇塞达Puigcerda 	feud.Barony
	苏埃尔特Suert 	feud.Barony
	特伦普Tremp 	feud.Barony
	乌尔赫尔Urgell 	feud.Barony
	巴利德沃伊Valledebohi 	feud.Barony
	别利亚Viella 	feud.Barony
}

func (c *乌尔赫尔UrgellCounty) BElpuidesegur厄尔贝塞居尔() feud.Barony {
	return c.厄尔贝塞居尔Elpuidesegur
}
    
func (c *乌尔赫尔UrgellCounty) BPallars帕拉尔斯() feud.Barony {
	return c.帕拉尔斯Pallars
}
    
func (c *乌尔赫尔UrgellCounty) BPuigcerda普奇塞达() feud.Barony {
	return c.普奇塞达Puigcerda
}
    
func (c *乌尔赫尔UrgellCounty) BSuert苏埃尔特() feud.Barony {
	return c.苏埃尔特Suert
}
    
func (c *乌尔赫尔UrgellCounty) BTremp特伦普() feud.Barony {
	return c.特伦普Tremp
}
    
func (c *乌尔赫尔UrgellCounty) BUrgell乌尔赫尔() feud.Barony {
	return c.乌尔赫尔Urgell
}
    
func (c *乌尔赫尔UrgellCounty) BValledebohi巴利德沃伊() feud.Barony {
	return c.巴利德沃伊Valledebohi
}
    
func (c *乌尔赫尔UrgellCounty) BViella别利亚() feud.Barony {
	return c.别利亚Viella
}
    
var CUrgell乌尔赫尔 UrgellCounty = &乌尔赫尔UrgellCounty{}

func init() {
	f := CUrgell乌尔赫尔.(*乌尔赫尔UrgellCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "206",
		Title:     "urgell",
		TitleName: "乌尔赫尔",
		TitleCode: "c_urgell",
		Baronies:  map[string]feud.Barony{},
	}

	f.厄尔贝塞居尔Elpuidesegur = BElpuidesegur厄尔贝塞居尔
	f.厄尔贝塞居尔Elpuidesegur.SetParent(f)
	
	f.帕拉尔斯Pallars = BPallars帕拉尔斯
	f.帕拉尔斯Pallars.SetParent(f)
	
	f.普奇塞达Puigcerda = BPuigcerda普奇塞达
	f.普奇塞达Puigcerda.SetParent(f)
	
	f.苏埃尔特Suert = BSuert苏埃尔特
	f.苏埃尔特Suert.SetParent(f)
	
	f.特伦普Tremp = BTremp特伦普
	f.特伦普Tremp.SetParent(f)
	
	f.乌尔赫尔Urgell = BUrgell乌尔赫尔
	f.乌尔赫尔Urgell.SetParent(f)
	
	f.巴利德沃伊Valledebohi = BValledebohi巴利德沃伊
	f.巴利德沃伊Valledebohi.SetParent(f)
	
	f.别利亚Viella = BViella别利亚
	f.别利亚Viella.SetParent(f)
	
}
