package talkerrepository

import (
	"encoding/json"
	"errors"
	"sort"
	"strings"

	talker "github.com/beto-ouverney/talker-manager-go/talker/entity"
)

//ITalkerRepository is the interface for the talker repository
type ITalkerRepository interface {
	GetAllTalkers() (*[]talker.Talker, error)
	GetTalkerByID(id int) (*talker.Talker, error)
	AddTalker(newTalker *talker.Talker) (*talker.Talker, error)
	EditTalker(newTalker *talker.Talker) (*talker.Talker, error)
	DeleteTalker(id int) error
	SearchTalkers(search string) (*[]talker.Talker, error)
}

//TalkerRepository is the implementation of the talker repository
type TalkerRepository struct {
	ITalkerRepository
}

//GetAllTalkers is a function that returns all the talkers
func (t *TalkerRepository) GetAllTalkers() (*[]talker.Talker, error) {
	jsonFile, err := readJSON()
	if err != nil {
		return nil, err
	}
	var talkers *[]talker.Talker
	err = json.Unmarshal(jsonFile, &talkers)
	return talkers, err
}

//GetTalkerByID is a function that returns a talker by id
func (t *TalkerRepository) GetTalkerByID(id int) (*talker.Talker, error) {
	jsonFile, err := readJSON()
	if err != nil {
		return nil, err
	}
	var talkers []talker.Talker
	err = json.Unmarshal(jsonFile, &talkers)
	res := talker.Talker{}
	var talker *talker.Talker
	for _, v := range talkers {
		if v.ID == id {
			res = v
			talker = &res
		}
	}
	return talker, err
}

// AddTalker adds a new talker to the list of talkers
func (t *TalkerRepository) AddTalker(newTalker *talker.Talker) (*talker.Talker, error) {
	jsonFile, err := readJSON()
	if err == nil {
		var talkers []talker.Talker
		err = json.Unmarshal(jsonFile, &talkers)
		if err == nil {
			sort.Slice(talkers, func(i, j int) bool { return talkers[i].ID < talkers[j].ID })
			lastTalker := talkers[len(talkers)-1]
			newTalker.ID = lastTalker.ID + 1
			talkers = append(talkers, *newTalker)
			jsonFile, err = json.MarshalIndent(talkers, "", "    ")
			if err == nil {
				err = writeJSON(jsonFile)
			}
		}
	}
	return newTalker, err
}

// EditTalker edits a talker
func (t *TalkerRepository) EditTalker(newTalker *talker.Talker) (*talker.Talker, error) {
	jsonFile, err := readJSON()
	if err == nil {
		var talkers []talker.Talker
		err = json.Unmarshal(jsonFile, &talkers)
		if err == nil {
			for i, v := range talkers {
				if v.ID == newTalker.ID {
					talkers[i] = *newTalker
				}
			}
			jsonFile, err = json.MarshalIndent(talkers, "", "    ")
			if err == nil {
				err = writeJSON(jsonFile)
			}
		}
	}
	return newTalker, err
}

//DeleteTalker deletes a talker from the list of talkers by id
func (t *TalkerRepository) DeleteTalker(id int) error {
	jsonFile, err := readJSON()
	if err == nil {
		var talkers []talker.Talker
		err = json.Unmarshal(jsonFile, &talkers)
		if err == nil {
			for i, v := range talkers {
				if v.ID == id {
					talkers = append(talkers[:i], talkers[i+1:]...)
					jsonFile, err = json.MarshalIndent(talkers, "", "    ")
					if err == nil {
						err = writeJSON(jsonFile)
						if err == nil {
							return err
						}
					}
				}
			}
		}
	}
	err = errors.New("Talker not found")
	return err
}

// SearchTalkers searches for talkers by name
func (t *TalkerRepository) SearchTalkers(search string) (*[]talker.Talker, error) {
	jsonFile, err := readJSON()
	if err == nil {
		var talkers []talker.Talker
		err = json.Unmarshal(jsonFile, &talkers)
		if search == "" {
			return &talkers, err
		}
		if err == nil {
			var talkersFound []talker.Talker
			for _, v := range talkers {
				str1 := strings.ToLower(v.Name)
				str2 := strings.ToLower(search)
				if strings.Contains(str1, str2) {
					talkersFound = append(talkersFound, v)
				}
			}
			if len(talkersFound) == 0 {
				talkersFound := make([]talker.Talker, 0)
				return &talkersFound, nil
			}
			return &talkersFound, nil
		}
	}
	return nil, err
}
