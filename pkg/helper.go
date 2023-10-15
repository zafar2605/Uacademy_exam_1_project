package pkg

// func ReadData(Path string) ([]any, error) {

// 	body, err := os.ReadFile(Path)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var data []any
// 	err = json.Unmarshal(body, &data)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return data, nil
// }

// func WriteData(Path string, data []any) error {

// 	body, err := json.MarshalIndent(data, "", "    ")
// 	if err != nil {
// 		return err
// 	}

// 	err = os.WriteFile(Path, body, os.ModePerm)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
