package tanggula

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TanggulaCounty interface {
    feud.County
    BDangla当拉() 	feud.Barony
    BMaqu玛曲() 	feud.Barony
    BMarong玛绒() 	feud.Barony
    BSewu色务() 	feud.Barony
    BTongtian通天() 	feud.Barony
    BWenquan温泉() 	feud.Barony
    BYanshiping雁石坪() 	feud.Barony
}

type 唐古拉TanggulaCounty struct {
	feud.BaseCounty
	当拉Dangla 	feud.Barony
	玛曲Maqu 	feud.Barony
	玛绒Marong 	feud.Barony
	色务Sewu 	feud.Barony
	通天Tongtian 	feud.Barony
	温泉Wenquan 	feud.Barony
	雁石坪Yanshiping 	feud.Barony
}

func (c *唐古拉TanggulaCounty) BDangla当拉() feud.Barony {
	return c.当拉Dangla
}
    
func (c *唐古拉TanggulaCounty) BMaqu玛曲() feud.Barony {
	return c.玛曲Maqu
}
    
func (c *唐古拉TanggulaCounty) BMarong玛绒() feud.Barony {
	return c.玛绒Marong
}
    
func (c *唐古拉TanggulaCounty) BSewu色务() feud.Barony {
	return c.色务Sewu
}
    
func (c *唐古拉TanggulaCounty) BTongtian通天() feud.Barony {
	return c.通天Tongtian
}
    
func (c *唐古拉TanggulaCounty) BWenquan温泉() feud.Barony {
	return c.温泉Wenquan
}
    
func (c *唐古拉TanggulaCounty) BYanshiping雁石坪() feud.Barony {
	return c.雁石坪Yanshiping
}
    
var CTanggula唐古拉 TanggulaCounty = &唐古拉TanggulaCounty{}

func init() {
	f := CTanggula唐古拉.(*唐古拉TanggulaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1573",
		Title:     "tanggula",
		TitleName: "唐古拉",
		TitleCode: "c_tanggula",
		Baronies:  map[string]feud.Barony{},
	}

	f.当拉Dangla = BDangla当拉
	f.当拉Dangla.SetParent(f)
	
	f.玛曲Maqu = BMaqu玛曲
	f.玛曲Maqu.SetParent(f)
	
	f.玛绒Marong = BMarong玛绒
	f.玛绒Marong.SetParent(f)
	
	f.色务Sewu = BSewu色务
	f.色务Sewu.SetParent(f)
	
	f.通天Tongtian = BTongtian通天
	f.通天Tongtian.SetParent(f)
	
	f.温泉Wenquan = BWenquan温泉
	f.温泉Wenquan.SetParent(f)
	
	f.雁石坪Yanshiping = BYanshiping雁石坪
	f.雁石坪Yanshiping.SetParent(f)
	
}
