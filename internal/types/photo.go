package types

import (
	"fmt"

	"github.com/theloosygoose/goserver/tools"
)

type CollectionChecked struct {
    Collection tools.Collection
    Value string
}

func (ck CollectionChecked) FillOff(collections []tools.Collection) []CollectionChecked {
    var c []CollectionChecked

    for _ ,collection := range collections{
        sc := CollectionChecked{
            Collection: collection,
            Value: "off",
        }
        c = append(c, sc)
    }
    return c
}

func (ck CollectionChecked) FillValues(collections []tools.Collection, values map[string]string) []CollectionChecked{
    var c []CollectionChecked

    for _, collection := range collections {
        cid := (collection.Name + fmt.Sprint(collection.ID))
        sc := CollectionChecked{
            Collection: collection,
            Value: values[cid],
        }
        c = append(c, sc)
    }

    return c
}

type FormValues struct{
    ID string
    Name string
    Location string
    Description string
    ImagePath string
    Collections []CollectionChecked
}

func FormwithValuesEmpty(collections []tools.Collection) FormValues {
    e := CollectionChecked{}.FillOff(collections)
        
    f := FormValues{
        ID: "",
        Name: "",
        Location: "",
        Description: "",
        ImagePath: "",
        Collections: e, 
    }

    return f
}

func FormwithValues(collections []tools.Collection, m map[string]string, photo tools.Photo) FormValues {
    e := CollectionChecked{}.FillValues(collections, m)

    f := FormValues{
        ID: fmt.Sprint(photo.ID),
        Name: photo.Name,
        Location: photo.Location,
        Description: photo.Description.String,
        ImagePath: photo.Imagepath,
        Collections: e,
    }
    return f
}
