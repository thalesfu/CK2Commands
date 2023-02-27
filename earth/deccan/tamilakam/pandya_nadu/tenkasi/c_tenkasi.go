package tenkasi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TenkasiCounty interface {
    feud.County
    BAivarmalai阿伐摩罗() 	feud.Barony
    BBandgaon槃陀伽罗摩() 	feud.Barony
    BDindigul丁迪古尔() 	feud.Barony
    BPalayamcottai波罗延柯多() 	feud.Barony
    BSankaranayinarkovil商羯罗那伊那罗拘毗罗() 	feud.Barony
    BTenkasi谛那迦尸() 	feud.Barony
    BTirumalapuram提卢摩罗补楞() 	feud.Barony
}

type 谛那迦尸TenkasiCounty struct {
	feud.BaseCounty
	阿伐摩罗Aivarmalai 	feud.Barony
	槃陀伽罗摩Bandgaon 	feud.Barony
	丁迪古尔Dindigul 	feud.Barony
	波罗延柯多Palayamcottai 	feud.Barony
	商羯罗那伊那罗拘毗罗Sankaranayinarkovil 	feud.Barony
	谛那迦尸Tenkasi 	feud.Barony
	提卢摩罗补楞Tirumalapuram 	feud.Barony
}

func (c *谛那迦尸TenkasiCounty) BAivarmalai阿伐摩罗() feud.Barony {
	return c.阿伐摩罗Aivarmalai
}
    
func (c *谛那迦尸TenkasiCounty) BBandgaon槃陀伽罗摩() feud.Barony {
	return c.槃陀伽罗摩Bandgaon
}
    
func (c *谛那迦尸TenkasiCounty) BDindigul丁迪古尔() feud.Barony {
	return c.丁迪古尔Dindigul
}
    
func (c *谛那迦尸TenkasiCounty) BPalayamcottai波罗延柯多() feud.Barony {
	return c.波罗延柯多Palayamcottai
}
    
func (c *谛那迦尸TenkasiCounty) BSankaranayinarkovil商羯罗那伊那罗拘毗罗() feud.Barony {
	return c.商羯罗那伊那罗拘毗罗Sankaranayinarkovil
}
    
func (c *谛那迦尸TenkasiCounty) BTenkasi谛那迦尸() feud.Barony {
	return c.谛那迦尸Tenkasi
}
    
func (c *谛那迦尸TenkasiCounty) BTirumalapuram提卢摩罗补楞() feud.Barony {
	return c.提卢摩罗补楞Tirumalapuram
}
    
var CTenkasi谛那迦尸 TenkasiCounty = &谛那迦尸TenkasiCounty{}

func init() {
	f := CTenkasi谛那迦尸.(*谛那迦尸TenkasiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1116",
		Title:     "tenkasi",
		TitleName: "谛那迦尸",
		TitleCode: "c_tenkasi",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿伐摩罗Aivarmalai = BAivarmalai阿伐摩罗
	f.阿伐摩罗Aivarmalai.SetParent(f)
	
	f.槃陀伽罗摩Bandgaon = BBandgaon槃陀伽罗摩
	f.槃陀伽罗摩Bandgaon.SetParent(f)
	
	f.丁迪古尔Dindigul = BDindigul丁迪古尔
	f.丁迪古尔Dindigul.SetParent(f)
	
	f.波罗延柯多Palayamcottai = BPalayamcottai波罗延柯多
	f.波罗延柯多Palayamcottai.SetParent(f)
	
	f.商羯罗那伊那罗拘毗罗Sankaranayinarkovil = BSankaranayinarkovil商羯罗那伊那罗拘毗罗
	f.商羯罗那伊那罗拘毗罗Sankaranayinarkovil.SetParent(f)
	
	f.谛那迦尸Tenkasi = BTenkasi谛那迦尸
	f.谛那迦尸Tenkasi.SetParent(f)
	
	f.提卢摩罗补楞Tirumalapuram = BTirumalapuram提卢摩罗补楞
	f.提卢摩罗补楞Tirumalapuram.SetParent(f)
	
}
