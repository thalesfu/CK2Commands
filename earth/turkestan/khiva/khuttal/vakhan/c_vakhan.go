package vakhan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VakhanCounty interface {
    feud.County
    BDaritubat吐蕃之门() 	feud.Barony
    BIshkashim塞迦审() 	feud.Barony
    BKhandud昏驮多() 	feud.Barony
    BKhorog霍罗格() 	feud.Barony
    BSast赛斯特() 	feud.Barony
    BSughnan识匿() 	feud.Barony
    BVakhan镬侃() 	feud.Barony
}

type 镬侃VakhanCounty struct {
	feud.BaseCounty
	吐蕃之门Daritubat 	feud.Barony
	塞迦审Ishkashim 	feud.Barony
	昏驮多Khandud 	feud.Barony
	霍罗格Khorog 	feud.Barony
	赛斯特Sast 	feud.Barony
	识匿Sughnan 	feud.Barony
	镬侃Vakhan 	feud.Barony
}

func (c *镬侃VakhanCounty) BDaritubat吐蕃之门() feud.Barony {
	return c.吐蕃之门Daritubat
}
    
func (c *镬侃VakhanCounty) BIshkashim塞迦审() feud.Barony {
	return c.塞迦审Ishkashim
}
    
func (c *镬侃VakhanCounty) BKhandud昏驮多() feud.Barony {
	return c.昏驮多Khandud
}
    
func (c *镬侃VakhanCounty) BKhorog霍罗格() feud.Barony {
	return c.霍罗格Khorog
}
    
func (c *镬侃VakhanCounty) BSast赛斯特() feud.Barony {
	return c.赛斯特Sast
}
    
func (c *镬侃VakhanCounty) BSughnan识匿() feud.Barony {
	return c.识匿Sughnan
}
    
func (c *镬侃VakhanCounty) BVakhan镬侃() feud.Barony {
	return c.镬侃Vakhan
}
    
var CVakhan镬侃 VakhanCounty = &镬侃VakhanCounty{}

func init() {
	f := CVakhan镬侃.(*镬侃VakhanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1543",
		Title:     "vakhan",
		TitleName: "镬侃",
		TitleCode: "c_vakhan",
		Baronies:  map[string]feud.Barony{},
	}

	f.吐蕃之门Daritubat = BDaritubat吐蕃之门
	f.吐蕃之门Daritubat.SetParent(f)
	
	f.塞迦审Ishkashim = BIshkashim塞迦审
	f.塞迦审Ishkashim.SetParent(f)
	
	f.昏驮多Khandud = BKhandud昏驮多
	f.昏驮多Khandud.SetParent(f)
	
	f.霍罗格Khorog = BKhorog霍罗格
	f.霍罗格Khorog.SetParent(f)
	
	f.赛斯特Sast = BSast赛斯特
	f.赛斯特Sast.SetParent(f)
	
	f.识匿Sughnan = BSughnan识匿
	f.识匿Sughnan.SetParent(f)
	
	f.镬侃Vakhan = BVakhan镬侃
	f.镬侃Vakhan.SetParent(f)
	
}
