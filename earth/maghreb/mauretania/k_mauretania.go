package mauretania

import (
	"github.com/thalesfu/CK2Commands/earth/maghreb/mauretania/adrar"
	"github.com/thalesfu/CK2Commands/earth/maghreb/mauretania/alger"
	"github.com/thalesfu/CK2Commands/earth/maghreb/mauretania/fes"
	"github.com/thalesfu/CK2Commands/earth/maghreb/mauretania/marrakech"
	"github.com/thalesfu/CK2Commands/earth/maghreb/mauretania/sijilmasa"
	"github.com/thalesfu/CK2Commands/earth/maghreb/mauretania/sous"
	"github.com/thalesfu/CK2Commands/earth/maghreb/mauretania/tahert"
	"github.com/thalesfu/CK2Commands/earth/maghreb/mauretania/tangiers"
	"github.com/thalesfu/CK2Commands/earth/maghreb/mauretania/tlemcen"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MauretaniaKingdom interface {
	feud.Kingdom
	DAdrar阿德拉尔() adrar.AdrarDuke
	DAlger阿尔及尔() alger.AlgerDuke
	DFes非斯() fes.FesDuke
	DMarrakech马拉喀什() marrakech.MarrakechDuke
	DSijilmasa锡吉勒马萨() sijilmasa.SijilmasaDuke
	DSous苏斯() sous.SousDuke
	DTahert提亚雷特() tahert.TahertDuke
	DTangiers丹吉尔() tangiers.TangiersDuke
	DTlemcen特莱姆森() tlemcen.TlemcenDuke
}

type 马格里布MauretaniaKingdom struct {
	feud.BaseKingdom
	阿德拉尔Adrar      adrar.AdrarDuke
	阿尔及尔Alger      alger.AlgerDuke
	非斯Fes          fes.FesDuke
	马拉喀什Marrakech  marrakech.MarrakechDuke
	锡吉勒马萨Sijilmasa sijilmasa.SijilmasaDuke
	苏斯Sous         sous.SousDuke
	提亚雷特Tahert     tahert.TahertDuke
	丹吉尔Tangiers    tangiers.TangiersDuke
	特莱姆森Tlemcen    tlemcen.TlemcenDuke
}

func (k *马格里布MauretaniaKingdom) DAdrar阿德拉尔() adrar.AdrarDuke {
	return k.阿德拉尔Adrar
}

func (k *马格里布MauretaniaKingdom) DAlger阿尔及尔() alger.AlgerDuke {
	return k.阿尔及尔Alger
}

func (k *马格里布MauretaniaKingdom) DFes非斯() fes.FesDuke {
	return k.非斯Fes
}

func (k *马格里布MauretaniaKingdom) DMarrakech马拉喀什() marrakech.MarrakechDuke {
	return k.马拉喀什Marrakech
}

func (k *马格里布MauretaniaKingdom) DSijilmasa锡吉勒马萨() sijilmasa.SijilmasaDuke {
	return k.锡吉勒马萨Sijilmasa
}

func (k *马格里布MauretaniaKingdom) DSous苏斯() sous.SousDuke {
	return k.苏斯Sous
}

func (k *马格里布MauretaniaKingdom) DTahert提亚雷特() tahert.TahertDuke {
	return k.提亚雷特Tahert
}

func (k *马格里布MauretaniaKingdom) DTangiers丹吉尔() tangiers.TangiersDuke {
	return k.丹吉尔Tangiers
}

func (k *马格里布MauretaniaKingdom) DTlemcen特莱姆森() tlemcen.TlemcenDuke {
	return k.特莱姆森Tlemcen
}

var KMauretania马格里布 MauretaniaKingdom = &马格里布MauretaniaKingdom{}

func init() {
	f := KMauretania马格里布.(*马格里布MauretaniaKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "mauretania",
		TitleName: "马格里布",
		TitleCode: "k_mauretania",
		Dukes:     map[string]feud.Duke{},
	}

	f.阿德拉尔Adrar = adrar.DAdrar阿德拉尔
	f.阿德拉尔Adrar.SetParent(f)

	f.阿尔及尔Alger = alger.DAlger阿尔及尔
	f.阿尔及尔Alger.SetParent(f)

	f.非斯Fes = fes.DFes非斯
	f.非斯Fes.SetParent(f)

	f.马拉喀什Marrakech = marrakech.DMarrakech马拉喀什
	f.马拉喀什Marrakech.SetParent(f)

	f.锡吉勒马萨Sijilmasa = sijilmasa.DSijilmasa锡吉勒马萨
	f.锡吉勒马萨Sijilmasa.SetParent(f)

	f.苏斯Sous = sous.DSous苏斯
	f.苏斯Sous.SetParent(f)

	f.提亚雷特Tahert = tahert.DTahert提亚雷特
	f.提亚雷特Tahert.SetParent(f)

	f.丹吉尔Tangiers = tangiers.DTangiers丹吉尔
	f.丹吉尔Tangiers.SetParent(f)

	f.特莱姆森Tlemcen = tlemcen.DTlemcen特莱姆森
	f.特莱姆森Tlemcen.SetParent(f)

}
