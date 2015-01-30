package interfaces

import "github.com/intercom/intercom-go/domain"

type EventAPI struct {
	httpClient HTTPClient
}

func NewEventAPI(httpClient HTTPClient) EventAPI {
	return EventAPI{httpClient: httpClient}
}

func (api EventAPI) Save(event domain.Event) error {
	_, err := api.httpClient.Post("/events", event)
	return err
}

//
// func parseResponseError(data []byte, statusCode int) HTTPError {
// 	errorList := HTTPErrorList{}
// 	err := json.Unmarshal(data, &errorList)
// 	if err != nil {
// 		return NewUnknownHTTPError()
// 	}
// 	return errorList.Errors[0] // only care about the first
// }
