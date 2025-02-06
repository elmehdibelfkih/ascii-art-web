package error

import (
	"html/template"
	"net/http"
)

type Error struct {
	StatusCode       int
	StatusText       string
	ErrorMessage     string
	ErrorTitle       string
	ErrorDescription string
}

var ErrorTmpl = template.Must(template.ParseFiles("./templates/error.html"))

func NotFoundError(w http.ResponseWriter, r *http.Request) {
	notFoundError := Error{
		StatusCode:       404,
		StatusText:       "Not Found",
		ErrorMessage:     "The page you are looking for might have been removed, had its name changed, or is temporarily unavailable.",
		ErrorTitle:       "Oops! Page Not Found",
		ErrorDescription: "We couldn't find the page you were looking for. Please check the URL for any mistakes or go back to the homepage.",
	}
	w.WriteHeader(notFoundError.StatusCode)
	ErrorTmpl.Execute(w, notFoundError)
}

func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	methodNotAllowed := Error{
		StatusCode:       405,
		StatusText:       "Method not allowed",
		ErrorMessage:     "The HTTP method used for this request is not allowed on this resource. Please use a different method.",
		ErrorTitle:       "Oops! method not allowed",
		ErrorDescription: "Please send a POST request to this endpoint with the required data for processing. Other HTTP methods are not allowed.",
	}
	w.WriteHeader(methodNotAllowed.StatusCode)
	ErrorTmpl.Execute(w, methodNotAllowed)
}

// func NotContainTheRequiredData(w http.ResponseWriter, r *http.Request) {
// 	notContainTheRequiredData := Error{
// 		StatusCode:       400,
// 		StatusText:       "Bad Request",
// 		ErrorMessage:     "The request is missing required data or is improperly formatted.",
// 		ErrorTitle:       "Oops! no data",
// 		ErrorDescription: "Please ensure you include all necessary fields in the request body and send a properly formatted JSON payload.",
// 	}
// 	w.WriteHeader(notContainTheRequiredData.StatusCode)
// 	ErrorTmpl.Execute(w, notContainTheRequiredData)
// }

func InternalServerError(w http.ResponseWriter, r *http.Request) {
	internalServerError := Error{
		StatusCode:       500,
		StatusText:       "Internal Server Error",
		ErrorMessage:     "An unexpected error occurred while processing your request.",
		ErrorTitle:       "Oops! Internal Server Error",
		ErrorDescription: "Something went wrong on our end. Please try again later or contact support if the issue persists.",
	}
	w.WriteHeader(internalServerError.StatusCode)
	ErrorTmpl.Execute(w, internalServerError)
}

func BadRequest(w http.ResponseWriter, r *http.Request) {
	badRequest := Error{
		StatusCode:       400,
		StatusText:       "Bad Request",
		ErrorMessage:     "The Server recieved a Bad request!!",
		ErrorTitle:       "Oops! Bad Request",
		ErrorDescription: "Please make sure you respect the input limits",
	}
	w.WriteHeader(badRequest.StatusCode)
	ErrorTmpl.Execute(w, badRequest)
}
