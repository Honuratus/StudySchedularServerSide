package people

import (
	"encoding/json"
	"strconv"
)

type Task struct {
	ID          int    `json:"id"`
	Context     string `json:"context"`
	IsCompleted int    `json:"is_completed"`
	ModelId     int    `json:"model_id"`
}

type Model struct {
	ID        int    `json:"id"`
	ModelName string `json:"model_name"`
	UserId    int    `json:"user_id"`
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (m *Model) UnmarshalJSON(data []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch v := raw["user_id"].(type) {
	case string:
		var err error
		m.UserId, err = strconv.Atoi(v)
		if err != nil {
			return err
		}
	}
	m.ModelName = raw["model_name"].(string)
	return nil
}

func (t *Task) UnmarshalJSON(data []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	var err error

	t.ModelId, err = strconv.Atoi(raw["model_id"].(string))
	if err != nil {
		return err
	}

	t.IsCompleted, err = strconv.Atoi(raw["is_completed"].(string))
	if err != nil {
		return err
	}
	t.Context = raw["context"].(string)

	return nil
}
