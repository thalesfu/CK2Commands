package simaramapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SimaramapuraCounty interface {
    feud.County
    BBettiah吠提阿() 	feud.Barony
    BBhaktagrama薄多伽罗摩() 	feud.Barony
    BKesaria计舍梨耶() 	feud.Barony
    BPattana跋多那() 	feud.Barony
    BSimaramapura僧摩罗摩城() 	feud.Barony
    BSkanderabad塞建咥罗跋陀() 	feud.Barony
    BVaid毗陀() 	feud.Barony
}

type 僧摩罗摩城SimaramapuraCounty struct {
	feud.BaseCounty
	吠提阿Bettiah 	feud.Barony
	薄多伽罗摩Bhaktagrama 	feud.Barony
	计舍梨耶Kesaria 	feud.Barony
	跋多那Pattana 	feud.Barony
	僧摩罗摩城Simaramapura 	feud.Barony
	塞建咥罗跋陀Skanderabad 	feud.Barony
	毗陀Vaid 	feud.Barony
}

func (c *僧摩罗摩城SimaramapuraCounty) BBettiah吠提阿() feud.Barony {
	return c.吠提阿Bettiah
}
    
func (c *僧摩罗摩城SimaramapuraCounty) BBhaktagrama薄多伽罗摩() feud.Barony {
	return c.薄多伽罗摩Bhaktagrama
}
    
func (c *僧摩罗摩城SimaramapuraCounty) BKesaria计舍梨耶() feud.Barony {
	return c.计舍梨耶Kesaria
}
    
func (c *僧摩罗摩城SimaramapuraCounty) BPattana跋多那() feud.Barony {
	return c.跋多那Pattana
}
    
func (c *僧摩罗摩城SimaramapuraCounty) BSimaramapura僧摩罗摩城() feud.Barony {
	return c.僧摩罗摩城Simaramapura
}
    
func (c *僧摩罗摩城SimaramapuraCounty) BSkanderabad塞建咥罗跋陀() feud.Barony {
	return c.塞建咥罗跋陀Skanderabad
}
    
func (c *僧摩罗摩城SimaramapuraCounty) BVaid毗陀() feud.Barony {
	return c.毗陀Vaid
}
    
var CSimaramapura僧摩罗摩城 SimaramapuraCounty = &僧摩罗摩城SimaramapuraCounty{}

func init() {
	f := CSimaramapura僧摩罗摩城.(*僧摩罗摩城SimaramapuraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1420",
		Title:     "simaramapura",
		TitleName: "僧摩罗摩城",
		TitleCode: "c_simaramapura",
		Baronies:  map[string]feud.Barony{},
	}

	f.吠提阿Bettiah = BBettiah吠提阿
	f.吠提阿Bettiah.SetParent(f)
	
	f.薄多伽罗摩Bhaktagrama = BBhaktagrama薄多伽罗摩
	f.薄多伽罗摩Bhaktagrama.SetParent(f)
	
	f.计舍梨耶Kesaria = BKesaria计舍梨耶
	f.计舍梨耶Kesaria.SetParent(f)
	
	f.跋多那Pattana = BPattana跋多那
	f.跋多那Pattana.SetParent(f)
	
	f.僧摩罗摩城Simaramapura = BSimaramapura僧摩罗摩城
	f.僧摩罗摩城Simaramapura.SetParent(f)
	
	f.塞建咥罗跋陀Skanderabad = BSkanderabad塞建咥罗跋陀
	f.塞建咥罗跋陀Skanderabad.SetParent(f)
	
	f.毗陀Vaid = BVaid毗陀
	f.毗陀Vaid.SetParent(f)
	
}
