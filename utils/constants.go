package utils

var StatusCodeToMessage = map[int]string{
	100: "Request accepted",
	200: "Success response with full object body",
	201: "Success response with Id only",
	400: "Invalid Input",
	404: "Record not found",
	424: "Failed to decode Body",
	500: "Error DB Fetch Failure / Server error",
}
