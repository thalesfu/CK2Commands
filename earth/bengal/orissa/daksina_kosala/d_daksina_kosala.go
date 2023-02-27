package daksina_kosala

import (
	"github.com/thalesfu/CK2Commands/earth/bengal/orissa/daksina_kosala/kiranapura"
	"github.com/thalesfu/CK2Commands/earth/bengal/orissa/daksina_kosala/rayapura"
	"github.com/thalesfu/CK2Commands/earth/bengal/orissa/daksina_kosala/sambalpur"
	"github.com/thalesfu/CK2Commands/earth/bengal/orissa/daksina_kosala/sripuri"
	"github.com/thalesfu/CK2Commands/earth/bengal/orissa/daksina_kosala/suvarnapura"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Daksina_kosalaDuke interface {
    feud.Duke
    CKiranapura日光城() 	kiranapura.KiranapuraCounty
    CRayapura罗耶补罗() 	rayapura.RayapuraCounty
    CSambalpur三婆罗补罗() 	sambalpur.SambalpurCounty
    CSripuri室利补梨() 	sripuri.SripuriCounty
    CSuvarnapura苏伐剌那补罗() 	suvarnapura.SuvarnapuraCounty
}

type 南憍萨罗Daksina_kosalaDuke struct {
	feud.BaseDuke
	日光城Kiranapura 	kiranapura.KiranapuraCounty
	罗耶补罗Rayapura 	rayapura.RayapuraCounty
	三婆罗补罗Sambalpur 	sambalpur.SambalpurCounty
	室利补梨Sripuri 	sripuri.SripuriCounty
	苏伐剌那补罗Suvarnapura 	suvarnapura.SuvarnapuraCounty
}

func (d *南憍萨罗Daksina_kosalaDuke) CKiranapura日光城() kiranapura.KiranapuraCounty {
	return d.日光城Kiranapura
}
    
func (d *南憍萨罗Daksina_kosalaDuke) CRayapura罗耶补罗() rayapura.RayapuraCounty {
	return d.罗耶补罗Rayapura
}
    
func (d *南憍萨罗Daksina_kosalaDuke) CSambalpur三婆罗补罗() sambalpur.SambalpurCounty {
	return d.三婆罗补罗Sambalpur
}
    
func (d *南憍萨罗Daksina_kosalaDuke) CSripuri室利补梨() sripuri.SripuriCounty {
	return d.室利补梨Sripuri
}
    
func (d *南憍萨罗Daksina_kosalaDuke) CSuvarnapura苏伐剌那补罗() suvarnapura.SuvarnapuraCounty {
	return d.苏伐剌那补罗Suvarnapura
}
    
var DDaksina_kosala南憍萨罗 Daksina_kosalaDuke = &南憍萨罗Daksina_kosalaDuke{}

func init() {
	f := DDaksina_kosala南憍萨罗.(*南憍萨罗Daksina_kosalaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "daksina_kosala",
		TitleName: "南憍萨罗",
		TitleCode: "d_daksina_kosala",
		Counties:  map[string]feud.County{},
	}

	f.日光城Kiranapura = kiranapura.CKiranapura日光城
	f.日光城Kiranapura.SetParent(f)
	
	f.罗耶补罗Rayapura = rayapura.CRayapura罗耶补罗
	f.罗耶补罗Rayapura.SetParent(f)
	
	f.三婆罗补罗Sambalpur = sambalpur.CSambalpur三婆罗补罗
	f.三婆罗补罗Sambalpur.SetParent(f)
	
	f.室利补梨Sripuri = sripuri.CSripuri室利补梨
	f.室利补梨Sripuri.SetParent(f)
	
	f.苏伐剌那补罗Suvarnapura = suvarnapura.CSuvarnapura苏伐剌那补罗
	f.苏伐剌那补罗Suvarnapura.SetParent(f)
	
}
