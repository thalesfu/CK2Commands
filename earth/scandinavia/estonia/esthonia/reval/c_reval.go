package reval

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RevalCounty interface {
    feud.County
    BBorpeal博雷阿尔() 	feud.Barony
    BHapsal哈普萨卢() 	feud.Barony
    BLaane莱内() 	feud.Barony
    BLeal利胡拉() 	feud.Barony
    BPades帕德斯() 	feud.Barony
    BParnaw派尔努() 	feud.Barony
    BReval卡莱万() 	feud.Barony
    BToompea座堂山() 	feud.Barony
}

type 卡莱万RevalCounty struct {
	feud.BaseCounty
	博雷阿尔Borpeal 	feud.Barony
	哈普萨卢Hapsal 	feud.Barony
	莱内Laane 	feud.Barony
	利胡拉Leal 	feud.Barony
	帕德斯Pades 	feud.Barony
	派尔努Parnaw 	feud.Barony
	卡莱万Reval 	feud.Barony
	座堂山Toompea 	feud.Barony
}

func (c *卡莱万RevalCounty) BBorpeal博雷阿尔() feud.Barony {
	return c.博雷阿尔Borpeal
}
    
func (c *卡莱万RevalCounty) BHapsal哈普萨卢() feud.Barony {
	return c.哈普萨卢Hapsal
}
    
func (c *卡莱万RevalCounty) BLaane莱内() feud.Barony {
	return c.莱内Laane
}
    
func (c *卡莱万RevalCounty) BLeal利胡拉() feud.Barony {
	return c.利胡拉Leal
}
    
func (c *卡莱万RevalCounty) BPades帕德斯() feud.Barony {
	return c.帕德斯Pades
}
    
func (c *卡莱万RevalCounty) BParnaw派尔努() feud.Barony {
	return c.派尔努Parnaw
}
    
func (c *卡莱万RevalCounty) BReval卡莱万() feud.Barony {
	return c.卡莱万Reval
}
    
func (c *卡莱万RevalCounty) BToompea座堂山() feud.Barony {
	return c.座堂山Toompea
}
    
var CReval卡莱万 RevalCounty = &卡莱万RevalCounty{}

func init() {
	f := CReval卡莱万.(*卡莱万RevalCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "378",
		Title:     "reval",
		TitleName: "卡莱万",
		TitleCode: "c_reval",
		Baronies:  map[string]feud.Barony{},
	}

	f.博雷阿尔Borpeal = BBorpeal博雷阿尔
	f.博雷阿尔Borpeal.SetParent(f)
	
	f.哈普萨卢Hapsal = BHapsal哈普萨卢
	f.哈普萨卢Hapsal.SetParent(f)
	
	f.莱内Laane = BLaane莱内
	f.莱内Laane.SetParent(f)
	
	f.利胡拉Leal = BLeal利胡拉
	f.利胡拉Leal.SetParent(f)
	
	f.帕德斯Pades = BPades帕德斯
	f.帕德斯Pades.SetParent(f)
	
	f.派尔努Parnaw = BParnaw派尔努
	f.派尔努Parnaw.SetParent(f)
	
	f.卡莱万Reval = BReval卡莱万
	f.卡莱万Reval.SetParent(f)
	
	f.座堂山Toompea = BToompea座堂山
	f.座堂山Toompea.SetParent(f)
	
}
