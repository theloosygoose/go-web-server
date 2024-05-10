// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.663
package admin

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
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
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"container w-full\"><form class=\"grid grid-cols-2 gap-6 align-top mt-12 *:block\" id=\"form\" hx-encoding=\"multipart/form-data\" hx-post=\"/addphoto\"><label class=\"col-start-1\" for=\"name-input\"><span>Name</span> <input class=\"w-full border-0 border-b-2 border-accent focus:ring-0\" type=\"text\" name=\"name\" id=\"name-input\"></label> <label class=\"col-start-1\" for=\"location-input\"><span>Location</span> <input class=\"w-full block border-0 border-b-2 border-accent focus:ring-0\" type=\"text\" name=\"location\" id=\"location-input\"></label> <label class=\"col-start-1\" for=\"description-input\"><span>Description</span> <textarea class=\"w-full block border-0 border-b-2 border-accent focus:ring-0\" rows=\"5\" name=\"description\" id=\"description-input\"></textarea></label> <label class=\"col-start-2 row-start-1 row-span-2\" for=\"imagefile-input\"><span>Image</span> <input class=\"w-full block border-0 border-accent focus:ring-0\n            file:focus:ring-0 file:border-0\n            file:px-5 file:py-2 file:rounded-sm file:bg-accent file:text-center\" type=\"file\" name=\"imageFile\" id=\"imagefile-input\"></label> <button class=\"w-fit h-fit px-5 py-2 rounded-sm col-start-1 bg-accent\n        focus:ring-0\n        text-center transition-all duration-150 hover:scale-105\n        inline-flex items-center\" type=\"submit\">Upload <svg class=\"htmx-indicator\" class=\"motion-reduce:hidden animate-spin ml-1 mr-3 h-5 w-5 text-white\" fill=\"none\" viewBox=\"0 0 24 24\"><circle class=\"opacity-25\" cx=\"12\" cy=\"12\" r=\"10\" stroke=\"currentColor\" stroke-width=\"4\"></circle> <path class=\"opacity-75\" fill=\"currentColor\" d=\"M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z\"></path></svg></button> <progress id=\"progress\" value=\"0\" max=\"100\"></progress></form></div><script type=\"text/javascript\">\n    htmx.on('#form', \"htmx:xhr:progress\", function (evt) {\n        htmx.find('#progress').setAttribute('value', evt.detail.loaded / evt.detail.total * 100)\n    })\n</script>")
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
