package lubelska

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LubelskaCounty interface {
    feud.County
    BGoraj戈拉伊() 	feud.Barony
    BLublin卢布林() 	feud.Barony
    BNadwisla下卡齐米日() 	feud.Barony
    BParczew帕尔切夫() 	feud.Barony
    BPodwieprza文奇纳() 	feud.Barony
    BUrzedow乌任杜夫() 	feud.Barony
    BZamosc扎莫希奇() 	feud.Barony
}

type 卢布林LubelskaCounty struct {
	feud.BaseCounty
	戈拉伊Goraj 	feud.Barony
	卢布林Lublin 	feud.Barony
	下卡齐米日Nadwisla 	feud.Barony
	帕尔切夫Parczew 	feud.Barony
	文奇纳Podwieprza 	feud.Barony
	乌任杜夫Urzedow 	feud.Barony
	扎莫希奇Zamosc 	feud.Barony
}

func (c *卢布林LubelskaCounty) BGoraj戈拉伊() feud.Barony {
	return c.戈拉伊Goraj
}
    
func (c *卢布林LubelskaCounty) BLublin卢布林() feud.Barony {
	return c.卢布林Lublin
}
    
func (c *卢布林LubelskaCounty) BNadwisla下卡齐米日() feud.Barony {
	return c.下卡齐米日Nadwisla
}
    
func (c *卢布林LubelskaCounty) BParczew帕尔切夫() feud.Barony {
	return c.帕尔切夫Parczew
}
    
func (c *卢布林LubelskaCounty) BPodwieprza文奇纳() feud.Barony {
	return c.文奇纳Podwieprza
}
    
func (c *卢布林LubelskaCounty) BUrzedow乌任杜夫() feud.Barony {
	return c.乌任杜夫Urzedow
}
    
func (c *卢布林LubelskaCounty) BZamosc扎莫希奇() feud.Barony {
	return c.扎莫希奇Zamosc
}
    
var CLubelska卢布林 LubelskaCounty = &卢布林LubelskaCounty{}

func init() {
	f := CLubelska卢布林.(*卢布林LubelskaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1585",
		Title:     "lubelska",
		TitleName: "卢布林",
		TitleCode: "c_lubelska",
		Baronies:  map[string]feud.Barony{},
	}

	f.戈拉伊Goraj = BGoraj戈拉伊
	f.戈拉伊Goraj.SetParent(f)
	
	f.卢布林Lublin = BLublin卢布林
	f.卢布林Lublin.SetParent(f)
	
	f.下卡齐米日Nadwisla = BNadwisla下卡齐米日
	f.下卡齐米日Nadwisla.SetParent(f)
	
	f.帕尔切夫Parczew = BParczew帕尔切夫
	f.帕尔切夫Parczew.SetParent(f)
	
	f.文奇纳Podwieprza = BPodwieprza文奇纳
	f.文奇纳Podwieprza.SetParent(f)
	
	f.乌任杜夫Urzedow = BUrzedow乌任杜夫
	f.乌任杜夫Urzedow.SetParent(f)
	
	f.扎莫希奇Zamosc = BZamosc扎莫希奇
	f.扎莫希奇Zamosc.SetParent(f)
	
}
