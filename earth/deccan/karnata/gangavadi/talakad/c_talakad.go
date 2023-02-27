package talakad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TalakadCounty interface {
    feud.County
    BKotagiri拘吒耆厘() 	feud.Barony
    BNanjangud难若那瞿荼() 	feud.Barony
    BNarasamangala那罗娑茫伽罗() 	feud.Barony
    BSivasamudram湿婆三慕达罗() 	feud.Barony
    BSomanathapura苏摩那他补罗() 	feud.Barony
    BTalakad多罗迦都() 	feud.Barony
    BUmmattur乌摩图尔() 	feud.Barony
    BVeppur吠补罗() 	feud.Barony
}

type 多罗迦都TalakadCounty struct {
	feud.BaseCounty
	拘吒耆厘Kotagiri 	feud.Barony
	难若那瞿荼Nanjangud 	feud.Barony
	那罗娑茫伽罗Narasamangala 	feud.Barony
	湿婆三慕达罗Sivasamudram 	feud.Barony
	苏摩那他补罗Somanathapura 	feud.Barony
	多罗迦都Talakad 	feud.Barony
	乌摩图尔Ummattur 	feud.Barony
	吠补罗Veppur 	feud.Barony
}

func (c *多罗迦都TalakadCounty) BKotagiri拘吒耆厘() feud.Barony {
	return c.拘吒耆厘Kotagiri
}
    
func (c *多罗迦都TalakadCounty) BNanjangud难若那瞿荼() feud.Barony {
	return c.难若那瞿荼Nanjangud
}
    
func (c *多罗迦都TalakadCounty) BNarasamangala那罗娑茫伽罗() feud.Barony {
	return c.那罗娑茫伽罗Narasamangala
}
    
func (c *多罗迦都TalakadCounty) BSivasamudram湿婆三慕达罗() feud.Barony {
	return c.湿婆三慕达罗Sivasamudram
}
    
func (c *多罗迦都TalakadCounty) BSomanathapura苏摩那他补罗() feud.Barony {
	return c.苏摩那他补罗Somanathapura
}
    
func (c *多罗迦都TalakadCounty) BTalakad多罗迦都() feud.Barony {
	return c.多罗迦都Talakad
}
    
func (c *多罗迦都TalakadCounty) BUmmattur乌摩图尔() feud.Barony {
	return c.乌摩图尔Ummattur
}
    
func (c *多罗迦都TalakadCounty) BVeppur吠补罗() feud.Barony {
	return c.吠补罗Veppur
}
    
var CTalakad多罗迦都 TalakadCounty = &多罗迦都TalakadCounty{}

func init() {
	f := CTalakad多罗迦都.(*多罗迦都TalakadCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1216",
		Title:     "talakad",
		TitleName: "多罗迦都",
		TitleCode: "c_talakad",
		Baronies:  map[string]feud.Barony{},
	}

	f.拘吒耆厘Kotagiri = BKotagiri拘吒耆厘
	f.拘吒耆厘Kotagiri.SetParent(f)
	
	f.难若那瞿荼Nanjangud = BNanjangud难若那瞿荼
	f.难若那瞿荼Nanjangud.SetParent(f)
	
	f.那罗娑茫伽罗Narasamangala = BNarasamangala那罗娑茫伽罗
	f.那罗娑茫伽罗Narasamangala.SetParent(f)
	
	f.湿婆三慕达罗Sivasamudram = BSivasamudram湿婆三慕达罗
	f.湿婆三慕达罗Sivasamudram.SetParent(f)
	
	f.苏摩那他补罗Somanathapura = BSomanathapura苏摩那他补罗
	f.苏摩那他补罗Somanathapura.SetParent(f)
	
	f.多罗迦都Talakad = BTalakad多罗迦都
	f.多罗迦都Talakad.SetParent(f)
	
	f.乌摩图尔Ummattur = BUmmattur乌摩图尔
	f.乌摩图尔Ummattur.SetParent(f)
	
	f.吠补罗Veppur = BVeppur吠补罗
	f.吠补罗Veppur.SetParent(f)
	
}
