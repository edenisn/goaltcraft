package goaltcraft

const databasesListsPath = "databases/list"

type DatabasesList struct {
	DatabasesList []Database `json:"data"`
}

type Database struct {
	Id     int      `json:"id"`
	Name   string   `json:"name"`
	UiTags []string `json:"ui_tags"`
	Groups []struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}
}

func (api *API) GetDatabasesList(fromId, limit int) (*DatabasesList, error) {
	type data struct {
		Token  string `json:"token"`
		FromId int    `json:"from_id"`
		Limit  int    `json:"limit"`
	}
	params := data{FromId: fromId, Limit: limit, Token: api.Token}

	databasesList := new(DatabasesList)
	_, err := api.Request("POST", databasesListsPath, params, &databasesList)
	if err != nil {
		return nil, err
	}

	return databasesList, nil
}
