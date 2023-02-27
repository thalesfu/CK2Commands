package nanded

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NandedCounty interface {
    feud.County
    BAurad奥罗陀() 	feud.Barony
    BBalol婆卢罗() 	feud.Barony
    BNanded难提陀() 	feud.Barony
    BParbhani帕尔博亚尼() 	feud.Barony
    BPathri帕特里() 	feud.Barony
    BQandhar坎达哈() 	feud.Barony
    BShelgaon势罗伽罗摩() 	feud.Barony
}

type 难提陀NandedCounty struct {
	feud.BaseCounty
	奥罗陀Aurad 	feud.Barony
	婆卢罗Balol 	feud.Barony
	难提陀Nanded 	feud.Barony
	帕尔博亚尼Parbhani 	feud.Barony
	帕特里Pathri 	feud.Barony
	坎达哈Qandhar 	feud.Barony
	势罗伽罗摩Shelgaon 	feud.Barony
}

func (c *难提陀NandedCounty) BAurad奥罗陀() feud.Barony {
	return c.奥罗陀Aurad
}
    
func (c *难提陀NandedCounty) BBalol婆卢罗() feud.Barony {
	return c.婆卢罗Balol
}
    
func (c *难提陀NandedCounty) BNanded难提陀() feud.Barony {
	return c.难提陀Nanded
}
    
func (c *难提陀NandedCounty) BParbhani帕尔博亚尼() feud.Barony {
	return c.帕尔博亚尼Parbhani
}
    
func (c *难提陀NandedCounty) BPathri帕特里() feud.Barony {
	return c.帕特里Pathri
}
    
func (c *难提陀NandedCounty) BQandhar坎达哈() feud.Barony {
	return c.坎达哈Qandhar
}
    
func (c *难提陀NandedCounty) BShelgaon势罗伽罗摩() feud.Barony {
	return c.势罗伽罗摩Shelgaon
}
    
var CNanded难提陀 NandedCounty = &难提陀NandedCounty{}

func init() {
	f := CNanded难提陀.(*难提陀NandedCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1258",
		Title:     "nanded",
		TitleName: "难提陀",
		TitleCode: "c_nanded",
		Baronies:  map[string]feud.Barony{},
	}

	f.奥罗陀Aurad = BAurad奥罗陀
	f.奥罗陀Aurad.SetParent(f)
	
	f.婆卢罗Balol = BBalol婆卢罗
	f.婆卢罗Balol.SetParent(f)
	
	f.难提陀Nanded = BNanded难提陀
	f.难提陀Nanded.SetParent(f)
	
	f.帕尔博亚尼Parbhani = BParbhani帕尔博亚尼
	f.帕尔博亚尼Parbhani.SetParent(f)
	
	f.帕特里Pathri = BPathri帕特里
	f.帕特里Pathri.SetParent(f)
	
	f.坎达哈Qandhar = BQandhar坎达哈
	f.坎达哈Qandhar.SetParent(f)
	
	f.势罗伽罗摩Shelgaon = BShelgaon势罗伽罗摩
	f.势罗伽罗摩Shelgaon.SetParent(f)
	
}
