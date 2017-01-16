package config

import
(
	"encoding/json"
	"os"
)

type Root struct
{
	SQL		MySQL	`json:"mysql"`
	Port	string	'json:"port"`
}

type MySQL struct
{
	Host	 string `json:"host"`
	Port	 string `json:"port"`
	User	 string `json:"user"`
	Password string	`json:"password"`
	Database string `json:"database"`
}

func Init(filename string) (*Root, error)
{
	file, err := os.Open(filename)
	if err != nil
	{
		return nil, err
	}
	
	decoder := json.NewDecoder(file)
	
	root := &Root{}
	err = decoder.Decode(root)
	if err != nil
	{
		return nil, err
	}
	
	return root, nil
}