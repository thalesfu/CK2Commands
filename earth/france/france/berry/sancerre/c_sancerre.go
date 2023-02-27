package sancerre

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SancerreCounty interface {
    feud.County
    BAssigny阿西尼() 	feud.Barony
    BCosne科讷() 	feud.Barony
    BCours库尔() 	feud.Barony
    BGien日安() 	feud.Barony
    BJars雅尔() 	feud.Barony
    BPouilly普伊() 	feud.Barony
    BSancerre桑塞尔() 	feud.Barony
}

type 桑塞尔SancerreCounty struct {
	feud.BaseCounty
	阿西尼Assigny 	feud.Barony
	科讷Cosne 	feud.Barony
	库尔Cours 	feud.Barony
	日安Gien 	feud.Barony
	雅尔Jars 	feud.Barony
	普伊Pouilly 	feud.Barony
	桑塞尔Sancerre 	feud.Barony
}

func (c *桑塞尔SancerreCounty) BAssigny阿西尼() feud.Barony {
	return c.阿西尼Assigny
}
    
func (c *桑塞尔SancerreCounty) BCosne科讷() feud.Barony {
	return c.科讷Cosne
}
    
func (c *桑塞尔SancerreCounty) BCours库尔() feud.Barony {
	return c.库尔Cours
}
    
func (c *桑塞尔SancerreCounty) BGien日安() feud.Barony {
	return c.日安Gien
}
    
func (c *桑塞尔SancerreCounty) BJars雅尔() feud.Barony {
	return c.雅尔Jars
}
    
func (c *桑塞尔SancerreCounty) BPouilly普伊() feud.Barony {
	return c.普伊Pouilly
}
    
func (c *桑塞尔SancerreCounty) BSancerre桑塞尔() feud.Barony {
	return c.桑塞尔Sancerre
}
    
var CSancerre桑塞尔 SancerreCounty = &桑塞尔SancerreCounty{}

func init() {
	f := CSancerre桑塞尔.(*桑塞尔SancerreCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1960",
		Title:     "sancerre",
		TitleName: "桑塞尔",
		TitleCode: "c_sancerre",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿西尼Assigny = BAssigny阿西尼
	f.阿西尼Assigny.SetParent(f)
	
	f.科讷Cosne = BCosne科讷
	f.科讷Cosne.SetParent(f)
	
	f.库尔Cours = BCours库尔
	f.库尔Cours.SetParent(f)
	
	f.日安Gien = BGien日安
	f.日安Gien.SetParent(f)
	
	f.雅尔Jars = BJars雅尔
	f.雅尔Jars.SetParent(f)
	
	f.普伊Pouilly = BPouilly普伊
	f.普伊Pouilly.SetParent(f)
	
	f.桑塞尔Sancerre = BSancerre桑塞尔
	f.桑塞尔Sancerre.SetParent(f)
	
}
