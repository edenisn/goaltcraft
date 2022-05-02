package altcraft

import (
	"encoding/json"
)

const databasesListsPath = "databases/list"

type ListOfDatabases struct {
	baseList
	DatabasesList []DatabasesListResponse `json:"data"`
}

type DatabasesListResponse struct {
	Id     int      `json:"id"`
	Name   string   `json:"name"`
	UiTags []string `json:"ui_tags"`
}

func (api *API) GetDatabasesList(limit int) (*ListOfDatabases, error) {
	data := new(ListOfDatabases)
	params := map[string]interface{}{"limit": limit}
	bytes, err := api.Request("POST", databasesListsPath, params)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
