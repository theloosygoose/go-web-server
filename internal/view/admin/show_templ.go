// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.663
package admin

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"fmt"
	"github.com/theloosygoose/goserver/internal/types"
	"os"

	"github.com/theloosygoose/goserver/internal/view/components"
	"github.com/theloosygoose/goserver/internal/view/layout"
)

func Show() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var2 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"container w-full\"><form class=\"grid grid-cols-2 gap-6 align-top mt-12 *:block\" id=\"form\" hx-encoding=\"multipart/form-data\" hx-target=\"#response-status\" hx-swap=\"innerHTML\" hx-put=\"/addphoto\" hx-on::after-request=\"this.reset()\"><label class=\"col-start-1\" for=\"name-input\"><span>Name</span> <input class=\"w-full border-0 border-b-2 border-accent focus:ring-0\n            invalid:border-pink-500 invalid:text-pink-600\" type=\"text\" name=\"name\" id=\"name-input\" required></label> <label class=\"col-start-1\" for=\"location-input\"><span>Location</span> <input class=\"w-full block border-0 border-b-2 border-accent focus:ring-0\" type=\"text\" name=\"location\" id=\"location-input\"></label> <label class=\"col-start-1\" for=\"description-input\"><span>Description</span> <textarea class=\"w-full block border-0 border-b-2 border-accent focus:ring-0\" rows=\"5\" name=\"description\" id=\"description-input\"></textarea></label> <label class=\"col-start-2 row-start-1 row-span-1\" for=\"imagefile-input\"><input class=\"w-full block border-0 border-accent focus:ring-0\n            file:transition-colors file:duration-150\n            file:focus:ring-0 file:border-0\n            file:hover:cursor-pointer\n            file:hover:bg-slate-700\n            file:px-5 file:py-2 file:rounded-sm \n            file:bg-slate-500 file:text-center\n            file:text-white\n            file:hover:invalid:bg-pink-700\n            file:invalid:border-pink-500 file:invalid:bg-pink-600\" type=\"file\" name=\"imageFile\" id=\"imagefile-input\" accept=\"image/*\" onchange=\"loadFile(event)\" required></label> <button class=\"w-fit h-fit px-5 py-2 rounded-sm col-start-1 bg-slate-500\n        relative focus:ring-0 text-center \n        transition-colors duration-150\n        text-white\n        hover:bg-slate-700\n        inline-flex items-center\" type=\"submit\">Upload <svg class=\"htmx-indicator relative right-0 bottom-6\n        animate-spin \n        ml-1 mr-3 h-5 w-5 text-white\" fill=\"none\" viewBox=\"0 0 12 12\"><circle class=\"opacity-25\" cx=\"12\" cy=\"12\" r=\"10\" stroke=\"currentColor\" stroke-width=\"4\"></circle> <path class=\"opacity-75\" fill=\"currentColor\" d=\"M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z\"></path></svg></button></form><div class=\"w-full flex justify-center\"><img class=\"max-h-56\" id=\"preview-image\" height=\"300px\"></div><div id=\"response-status\"></div></div><button class=\"w-fit h-fit px-5 py-2 rounded-sm col-start-1 bg-slate-500\n        my-16\n        relative focus:ring-0 text-center \n        transition-colors duration-150\n        justify-self-center\n        text-white\n        hover:bg-slate-700\n        inline-flex items-center\" hx-get=\"/photodata/delete\" hx-swap=\"innerHTML\" hx-target=\"#delete-items\" hx-trigger=\"click\">Delete Photos </button><div id=\"delete-items\"></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = layout.Base().Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func Delete(photos []types.Photo) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var3 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var3 == nil {
			templ_7745c5c3_Var3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var4 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			for _, photo := range photos {
				var templ_7745c5c3_Var5 = []any{"group hover:cursor-pointer relative blur-load z-0 hover:z-10 w-fit h-fit hover:scale-105 ease-in-out duration-200"}
				templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var5...)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var6 string
				templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var5).String())
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/view/admin/show.templ`, Line: 1, Col: 0}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-trigger=\"click\" hx-delete=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var7 string
				templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("photodata/%d", photo.ID))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/view/admin/show.templ`, Line: 113, Col: 65}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"this\" hx-swap=\"delete\" hx-confirm=\"Are you sure you want to delete this photo?\"><img class=\"z-20 object-center object-cover max-w-44 max-h-44 \n                    group-hover:blur-sm\n                    transition-all duration-300\" loading=\"lazy\" src=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var8 string
				templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(os.Getenv("PHOTO_DIR") + "med_" + photo.Image.FileName)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/view/admin/show.templ`, Line: 123, Col: 80}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" alt=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var9 string
				templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(photo.Name)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/view/admin/show.templ`, Line: 124, Col: 36}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><div class=\"absolute inset-0 bottom-0 left-0\n                            w-full h-full text-white text-center font-bold text-lg\n                            opacity-0 group-hover:opacity-50 bg-red-700 \n                            transition-all duration-300 ease-in-out\"></div></div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = components.Gallery().Render(templ.WithChildren(ctx, templ_7745c5c3_Var4), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
