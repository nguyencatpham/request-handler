package requestHandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/labstack/echo"
)

// TokenType type of token
type TokenType int8

// const of token type
const (
	Basic = iota
	Bearer
)

func (s TokenType) String() string {
	switch s {
	case Basic:
		return "Basic"
	case Bearer:
		return "Bearer"
	default:
		return "Bearer"
	}
	return "Bearer"
}

// RequestModel model
type RequestModel struct {
	URL       string
	TokenType TokenType
	Token     string
	Username  string
	Password  string
	Body      string
}

// Get request
func Get(requestModel RequestModel, result *interface{}) error {
	req, err := http.NewRequest("GET", requestModel.URL, nil)
	switch requestModel.TokenType {
	case Basic:
		req.SetBasicAuth(requestModel.Username, requestModel.Password)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	case Bearer:
		header := fmt.Sprintf("Bearer %s", requestModel.Token)
		req.Header.Set("Authorization", header)
	default:
		break
	}
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(result); err != nil {
		fmt.Printf("GET-URL:%s-ERROR:%v\n", requestModel.URL, err)
		return nil
	}
	if resp.StatusCode >= http.StatusBadRequest {
		message := fmt.Sprintf("GET-URL:%s-ERROR CODE:%v\n", requestModel.URL, resp.StatusCode)
		if result != nil && !reflect.ValueOf(result).IsNil() {
			temp, ok := (*result).(map[string]interface{})
			if ok && temp != nil && temp["data"] != nil {
				data := (temp["data"]).(map[string]interface{})
				if data != nil {
					msg := data["message"]
					if msg != nil {
						msgStr := msg.(string)

						if msgStr != "" {
							message = msgStr
						}
					}
				}
			}
		}
		return echo.NewHTTPError(resp.StatusCode, message)
	}

	return nil
}

// Post request
func Post(requestModel RequestModel, result *interface{}) error {
	req, err := http.NewRequest("POST", requestModel.URL, nil)
	switch requestModel.TokenType {
	case Basic:
		req.SetBasicAuth(requestModel.Username, requestModel.Password)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	case Bearer:
		header := fmt.Sprintf("Bearer %s", requestModel.Token)
		req.Header.Set("Authorization", header)
	default:
		break
	}
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(result); err != nil {
		fmt.Printf("POST-URL:%s-ERROR:%v\n", requestModel.URL, err)
		return nil
	}
	if resp.StatusCode >= http.StatusBadRequest {
		message := fmt.Sprintf("POST-URL:%s-ERROR CODE:%v\n", requestModel.URL, resp.StatusCode)
		if result != nil && !reflect.ValueOf(result).IsNil() {
			temp, ok := (*result).(map[string]interface{})
			if ok && temp != nil && temp["data"] != nil {
				data := (temp["data"]).(map[string]interface{})
				if data != nil {
					msg := data["message"]
					if msg != nil {
						msgStr := msg.(string)

						if msgStr != "" {
							message = msgStr
						}
					}
				}
			}
		}
		return echo.NewHTTPError(resp.StatusCode, message)
	}

	return nil
}

// Put request
func Put(requestModel RequestModel, result *interface{}) error {
	req, err := http.NewRequest("PUT", requestModel.URL, nil)
	switch requestModel.TokenType {
	case Basic:
		req.SetBasicAuth(requestModel.Username, requestModel.Password)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	case Bearer:
		header := fmt.Sprintf("Bearer %s", requestModel.Token)
		req.Header.Set("Authorization", header)
	default:
		break
	}
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(result); err != nil {
		fmt.Printf("PUT-URL:%s-ERROR:%v\n", requestModel.URL, err)
		return nil
	}
	if resp.StatusCode >= http.StatusBadRequest {
		message := fmt.Sprintf("PUT-URL:%s-ERROR CODE:%v\n", requestModel.URL, resp.StatusCode)
		if result != nil && !reflect.ValueOf(result).IsNil() {
			temp, ok := (*result).(map[string]interface{})
			if ok && temp != nil && temp["data"] != nil {
				data := (temp["data"]).(map[string]interface{})
				if data != nil {
					msg := data["message"]
					if msg != nil {
						msgStr := msg.(string)

						if msgStr != "" {
							message = msgStr
						}
					}
				}
			}
		}
		return echo.NewHTTPError(resp.StatusCode, message)
	}

	return nil
}

// Patch request
func Patch(requestModel RequestModel, result *interface{}) error {
	req, err := http.NewRequest("PATCH", requestModel.URL, nil)
	switch requestModel.TokenType {
	case Basic:
		req.SetBasicAuth(requestModel.Username, requestModel.Password)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	case Bearer:
		header := fmt.Sprintf("Bearer %s", requestModel.Token)
		req.Header.Set("Authorization", header)
	default:
		break
	}
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(result); err != nil {
		fmt.Printf("PATCH-URL:%s-ERROR:%v\n", requestModel.URL, err)
		return nil
	}
	if resp.StatusCode >= http.StatusBadRequest {
		message := fmt.Sprintf("PATCH-URL:%s-ERROR CODE:%v\n", requestModel.URL, resp.StatusCode)
		if result != nil && !reflect.ValueOf(result).IsNil() {
			temp, ok := (*result).(map[string]interface{})
			if ok && temp != nil && temp["data"] != nil {
				data := (temp["data"]).(map[string]interface{})
				if data != nil {
					msg := data["message"]
					if msg != nil {
						msgStr := msg.(string)

						if msgStr != "" {
							message = msgStr
						}
					}
				}
			}
		}
		return echo.NewHTTPError(resp.StatusCode, message)
	}

	return nil
}

// Delete request
func Delete(requestModel RequestModel, result *interface{}) error {
	req, err := http.NewRequest("DELETE", requestModel.URL, nil)
	switch requestModel.TokenType {
	case Basic:
		req.SetBasicAuth(requestModel.Username, requestModel.Password)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	case Bearer:
		header := fmt.Sprintf("Bearer %s", requestModel.Token)
		req.Header.Set("Authorization", header)
	default:
		break
	}
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(result); err != nil {
		fmt.Printf("DELETE-URL:%s-ERROR:%v\n", requestModel.URL, err)
		return nil
	}
	if resp.StatusCode >= http.StatusBadRequest {
		message := fmt.Sprintf("DELETE-URL:%s-ERROR CODE:%v\n", requestModel.URL, resp.StatusCode)
		if result != nil && !reflect.ValueOf(result).IsNil() {
			temp, ok := (*result).(map[string]interface{})
			if ok && temp != nil && temp["data"] != nil {
				data := (temp["data"]).(map[string]interface{})
				if data != nil {
					msg := data["message"]
					if msg != nil {
						msgStr := msg.(string)

						if msgStr != "" {
							message = msgStr
						}
					}
				}
			}
		}
		return echo.NewHTTPError(resp.StatusCode, message)
	}

	return nil
}
