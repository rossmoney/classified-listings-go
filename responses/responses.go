package responses

// Model used for form field validation errors
// swagger:model ValidationErrorResponse
type ValidationErrorResponse struct {
	//	@ApiModelProperty( value = "array of form field validation errors",
	//            name = "errors",
	//            dataType = "array",
	//            required = true,
	//            position = 1
	//   )
	Errors []Error `json:"errors"`
	//	@ApiModelProperty( value = "http status",
	//            name = "status",
	//            dataType = "int",
	//            position = 1
	//   )
	Status int `json:"status"`
}

// Model used for failure with error strings
// swagger:model ErrorMessageResponse
type ErrorMessageResponse struct {
	//	@ApiModelProperty( value = "error message",
	//            name = "message",
	//            dataType = "string",
	//            required = true,
	//            position = 1
	//   )
	Message string `json:"message"`
	//	@ApiModelProperty( value = "error output",
	//            name = "error",
	//            dataType = "string",
	//            required = true,
	//            position = 2
	//   )
	Error string `json:"error"`
	//	@ApiModelProperty( value = "http status",
	//            name = "status",
	//            dataType = "int",
	//            position = 3
	//   )
	Status int `json:"status"`
}

// Model used for success with data returned
// swagger:model SuccessResponse
type SuccessResponse struct {
	//	@ApiModelProperty( value = "message to display",
	//            name = "message",
	//            dataType = "string",
	//            position = 1
	//   )
	Message string `json:"message"`
	//	@ApiModelProperty( value = "http status",
	//            name = "status",
	//            dataType = "int",
	//            position = 2
	//   )
	Status int `json:"status"`
	//	@ApiModelProperty( value = "data",
	//            name = "data",
	//            dataType = "array",
	//            required = true,
	//            position = 3
	//   )
	Data any `json:"data"`
}

type Error struct {
	Message   string `json:"message"`
	FieldName string `json:"field_name"`
}
