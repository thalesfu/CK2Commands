package tus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TusCounty interface {
    feud.County
    BBojnurd博季努尔德() 	feud.Barony
    BFagdatdezh法达德次() 	feud.Barony
    BGhaisar盖萨尔() 	feud.Barony
    BHasanabad哈桑阿巴德() 	feud.Barony
    BIsfarayen埃斯法拉延() 	feud.Barony
    BMashhad马什哈德() 	feud.Barony
    BShervan希尔万() 	feud.Barony
    BSolak索拉克() 	feud.Barony
    BTus图斯() 	feud.Barony
}

type 图斯TusCounty struct {
	feud.BaseCounty
	博季努尔德Bojnurd 	feud.Barony
	法达德次Fagdatdezh 	feud.Barony
	盖萨尔Ghaisar 	feud.Barony
	哈桑阿巴德Hasanabad 	feud.Barony
	埃斯法拉延Isfarayen 	feud.Barony
	马什哈德Mashhad 	feud.Barony
	希尔万Shervan 	feud.Barony
	索拉克Solak 	feud.Barony
	图斯Tus 	feud.Barony
}

func (c *图斯TusCounty) BBojnurd博季努尔德() feud.Barony {
	return c.博季努尔德Bojnurd
}
    
func (c *图斯TusCounty) BFagdatdezh法达德次() feud.Barony {
	return c.法达德次Fagdatdezh
}
    
func (c *图斯TusCounty) BGhaisar盖萨尔() feud.Barony {
	return c.盖萨尔Ghaisar
}
    
func (c *图斯TusCounty) BHasanabad哈桑阿巴德() feud.Barony {
	return c.哈桑阿巴德Hasanabad
}
    
func (c *图斯TusCounty) BIsfarayen埃斯法拉延() feud.Barony {
	return c.埃斯法拉延Isfarayen
}
    
func (c *图斯TusCounty) BMashhad马什哈德() feud.Barony {
	return c.马什哈德Mashhad
}
    
func (c *图斯TusCounty) BShervan希尔万() feud.Barony {
	return c.希尔万Shervan
}
    
func (c *图斯TusCounty) BSolak索拉克() feud.Barony {
	return c.索拉克Solak
}
    
func (c *图斯TusCounty) BTus图斯() feud.Barony {
	return c.图斯Tus
}
    
var CTus图斯 TusCounty = &图斯TusCounty{}

func init() {
	f := CTus图斯.(*图斯TusCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "632",
		Title:     "tus",
		TitleName: "图斯",
		TitleCode: "c_tus",
		Baronies:  map[string]feud.Barony{},
	}

	f.博季努尔德Bojnurd = BBojnurd博季努尔德
	f.博季努尔德Bojnurd.SetParent(f)
	
	f.法达德次Fagdatdezh = BFagdatdezh法达德次
	f.法达德次Fagdatdezh.SetParent(f)
	
	f.盖萨尔Ghaisar = BGhaisar盖萨尔
	f.盖萨尔Ghaisar.SetParent(f)
	
	f.哈桑阿巴德Hasanabad = BHasanabad哈桑阿巴德
	f.哈桑阿巴德Hasanabad.SetParent(f)
	
	f.埃斯法拉延Isfarayen = BIsfarayen埃斯法拉延
	f.埃斯法拉延Isfarayen.SetParent(f)
	
	f.马什哈德Mashhad = BMashhad马什哈德
	f.马什哈德Mashhad.SetParent(f)
	
	f.希尔万Shervan = BShervan希尔万
	f.希尔万Shervan.SetParent(f)
	
	f.索拉克Solak = BSolak索拉克
	f.索拉克Solak.SetParent(f)
	
	f.图斯Tus = BTus图斯
	f.图斯Tus.SetParent(f)
	
}
