package ashir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AshirCounty interface {
    feud.County
    BAsir_banya阿西尔_班亚() 	feud.Barony
    BAsir_yasir阿西尔_亚西尔() 	feud.Barony
    BIboutarene伊布塔赫恩() 	feud.Barony
    BMouddat穆达特() 	feud.Barony
    BRoumana鲁马纳() 	feud.Barony
    BSidiheg西迪赫格() 	feud.Barony
    BSidinhar西迪恩哈尔() 	feud.Barony
}

type 阿希尔AshirCounty struct {
	feud.BaseCounty
	阿西尔_班亚Asir_banya 	feud.Barony
	阿西尔_亚西尔Asir_yasir 	feud.Barony
	伊布塔赫恩Iboutarene 	feud.Barony
	穆达特Mouddat 	feud.Barony
	鲁马纳Roumana 	feud.Barony
	西迪赫格Sidiheg 	feud.Barony
	西迪恩哈尔Sidinhar 	feud.Barony
}

func (c *阿希尔AshirCounty) BAsir_banya阿西尔_班亚() feud.Barony {
	return c.阿西尔_班亚Asir_banya
}
    
func (c *阿希尔AshirCounty) BAsir_yasir阿西尔_亚西尔() feud.Barony {
	return c.阿西尔_亚西尔Asir_yasir
}
    
func (c *阿希尔AshirCounty) BIboutarene伊布塔赫恩() feud.Barony {
	return c.伊布塔赫恩Iboutarene
}
    
func (c *阿希尔AshirCounty) BMouddat穆达特() feud.Barony {
	return c.穆达特Mouddat
}
    
func (c *阿希尔AshirCounty) BRoumana鲁马纳() feud.Barony {
	return c.鲁马纳Roumana
}
    
func (c *阿希尔AshirCounty) BSidiheg西迪赫格() feud.Barony {
	return c.西迪赫格Sidiheg
}
    
func (c *阿希尔AshirCounty) BSidinhar西迪恩哈尔() feud.Barony {
	return c.西迪恩哈尔Sidinhar
}
    
var CAshir阿希尔 AshirCounty = &阿希尔AshirCounty{}

func init() {
	f := CAshir阿希尔.(*阿希尔AshirCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1734",
		Title:     "ashir",
		TitleName: "阿希尔",
		TitleCode: "c_ashir",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿西尔_班亚Asir_banya = BAsir_banya阿西尔_班亚
	f.阿西尔_班亚Asir_banya.SetParent(f)
	
	f.阿西尔_亚西尔Asir_yasir = BAsir_yasir阿西尔_亚西尔
	f.阿西尔_亚西尔Asir_yasir.SetParent(f)
	
	f.伊布塔赫恩Iboutarene = BIboutarene伊布塔赫恩
	f.伊布塔赫恩Iboutarene.SetParent(f)
	
	f.穆达特Mouddat = BMouddat穆达特
	f.穆达特Mouddat.SetParent(f)
	
	f.鲁马纳Roumana = BRoumana鲁马纳
	f.鲁马纳Roumana.SetParent(f)
	
	f.西迪赫格Sidiheg = BSidiheg西迪赫格
	f.西迪赫格Sidiheg.SetParent(f)
	
	f.西迪恩哈尔Sidinhar = BSidinhar西迪恩哈尔
	f.西迪恩哈尔Sidinhar.SetParent(f)
	
}
