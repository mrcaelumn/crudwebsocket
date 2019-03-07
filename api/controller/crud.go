package controller

import (
	"encoding/json"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/goadesign/goa"
	"github.com/inconshreveable/log15"
	"github.com/mrcaelumn/crudwebsocket/api/app"
	"github.com/mrcaelumn/crudwebsocket/global"
)

// CrudController implements the crud resource.
type CrudController struct {
	*goa.Controller
}

// NewCrudController creates a crud controller.
func NewCrudController(service *goa.Service) *CrudController {
	return &CrudController{Controller: service.NewController("CrudController")}
}

// Create runs the create action.
func (c *CrudController) Create(ctx *app.CreateCrudContext) error {
	// CrudController_Create: start_implement

	// Put your logic here
	var objmap map[string]string
	decoder := json.NewDecoder(ctx.Body)
	defer ctx.Body.Close()
	err := decoder.Decode(&objmap)
	if err != nil {
		log15.Error(err.Error())
		return ctx.BadRequest(&app.CrudwebsocketError{Msg: "failed to parse json", Code: "002"})
	}
	if objmap["id"] == "" {
		return ctx.BadRequest(&app.CrudwebsocketError{Msg: "ID not found", Code: "003"})
	}
	if objmap["text"] == "" {
		return ctx.BadRequest(&app.CrudwebsocketError{Msg: "text not found", Code: "004"})
	}
	curTime := strconv.Itoa(int(time.Now().Unix()))
	text := make(map[string]string)

	text[curTime] = objmap["text"]

	global.DATAS[objmap["id"]] = append(global.DATAS[objmap["id"]], text)
	return ctx.OK([]byte(`{ "status" : "updated" }`))
	// CrudController_Create: end_implement
}

// Delete runs the delete action.
func (c *CrudController) Delete(ctx *app.DeleteCrudContext) error {
	// CrudController_Delete: start_implement

	// Put your logic here
	var objmap map[string]string
	decoder := json.NewDecoder(ctx.Body)
	defer ctx.Body.Close()
	err := decoder.Decode(&objmap)
	if ctx.ID == "" {
		return ctx.BadRequest(&app.CrudwebsocketError{Msg: "ID not found", Code: "003"})
	}
	if err != nil {
		log15.Error(err.Error())
		return ctx.BadRequest(&app.CrudwebsocketError{Msg: "failed to parse json", Code: "002"})
	}
	if objmap["idchat"] == "" {
		return ctx.BadRequest(&app.CrudwebsocketError{Msg: "ID CHAT not found", Code: "003"})
	}
	if objmap["text"] == "" {
		return ctx.BadRequest(&app.CrudwebsocketError{Msg: "text not found", Code: "004"})
	}

	data := global.DATAS[strings.ToLower(strings.TrimSpace(ctx.ID))]
	updated := make([]map[string]string, len(data))
	for _, val := range data {
		_, ok := val[strings.ToLower(strings.TrimSpace(objmap["idchat"]))]
		if !ok {
			updated = append(updated, val)
		}
	}
	global.DATAS[strings.ToLower(strings.TrimSpace(ctx.ID))] = updated

	return ctx.OK([]byte(`{ "status" : "` + objmap["idchat"] + ` deleted" }`))
	// CrudController_Delete: end_implement
}

type message struct {
	User string `json:"user"`
	Date int    `json:"date"`
	Text string `json:"text"`
}

// Select runs the select action.
func (c *CrudController) Select(ctx *app.SelectCrudContext) error {
	// CrudController_Select: start_implement

	// Put your logic here

	if strings.ToLower(strings.TrimSpace(ctx.ID)) == "all" {
		jsonString, err := json.Marshal(global.DATAS)
		if err != nil {
			return ctx.BadRequest(&app.CrudwebsocketError{Msg: err.Error(), Code: "001"})
		}
		return ctx.OK(jsonString)
	} else if strings.ToLower(strings.TrimSpace(ctx.ID)) == "sorted" {
		// messages := make([]message, 100)
		var messages []message
		for key, val := range global.DATAS {
			for _, va := range val {
				for k, v := range va {
					d, _ := strconv.Atoi(k)
					messages = append(messages, message{
						User: key,
						Date: d,
						Text: v,
					})
				}
			}
		}
		sort.Slice(messages, func(i, j int) bool {
			return messages[i].Date > messages[j].Date
		})

		jsonString, err := json.Marshal(messages)
		if err != nil {
			return ctx.BadRequest(&app.CrudwebsocketError{Msg: err.Error(), Code: "001"})
		}
		return ctx.OK(jsonString)
	} else {
		data := global.DATAS[strings.ToLower(strings.TrimSpace(ctx.ID))]

		jsonString, err := json.Marshal(data)
		if err != nil {
			return ctx.BadRequest(&app.CrudwebsocketError{Msg: err.Error(), Code: "001"})
		}
		return ctx.OK(jsonString)
	}

	// CrudController_Select: end_implement
}

// Update runs the update action.
func (c *CrudController) Update(ctx *app.UpdateCrudContext) error {
	// CrudController_Update: start_implement

	// Put your logic here
	var objmap map[string]string
	decoder := json.NewDecoder(ctx.Body)
	defer ctx.Body.Close()
	err := decoder.Decode(&objmap)
	if err != nil {
		log15.Error(err.Error())
		return ctx.BadRequest(&app.CrudwebsocketError{Msg: "failed to parse json", Code: "002"})
	}
	if objmap["idchat"] == "" {
		return ctx.BadRequest(&app.CrudwebsocketError{Msg: "ID CHAT not found", Code: "003"})
	}
	if objmap["text"] == "" {
		return ctx.BadRequest(&app.CrudwebsocketError{Msg: "text not found", Code: "004"})
	}

	data := global.DATAS[strings.ToLower(strings.TrimSpace(ctx.ID))]
	updated := make([]map[string]string, len(data))
	for _, val := range data {
		_, ok := val[strings.ToLower(strings.TrimSpace(objmap["idchat"]))]
		if ok {
			val[strings.ToLower(strings.TrimSpace(objmap["idchat"]))] = objmap["text"]
			updated = append(updated, val)
		} else {
			updated = append(updated, val)
		}
	}
	global.DATAS[strings.ToLower(strings.TrimSpace(ctx.ID))] = updated
	return ctx.OK([]byte(`{ "status" : "` + objmap["idchat"] + ` updated" }`))
	// CrudController_Update: end_implement
}

func delete_empty(s []message) []message {
	var r []message
	for _, str := range s {
		if str.Text != "" {
			r = append(r, str)
		}
	}
	return r
}
